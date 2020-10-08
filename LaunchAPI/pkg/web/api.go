package webapi

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"

	constants "epyphite/space/v1/LaunchAPI/pkg/constants"
	models "epyphite/space/v1/LaunchAPI/pkg/models"
	"epyphite/space/v1/LaunchAPI/pkg/storage"
	webapi "epyphite/space/v1/LaunchAPI/pkg/web/webapp"
)

//APIOne main structure
type APIOne struct {
	webConfig models.Config
}

//NewWebAgent // creates a mew instace \of web one
func NewWebAgent(config models.Config) (APIOne, error) {
	var APIOne APIOne
	APIOne.webConfig = config
	return APIOne, nil
}

//StartServer Starts the server using the variable sip and port, creates anew instance.
func (W *APIOne) StartServer() {
	log.Infoln("Version : " + constants.BuildVersion)
	log.Infoln("Build Time : " + constants.BuildTime)
	handler := W.New()
	log.Infof("Starting Server at %s:%s \n", W.webConfig.WebAddress, W.webConfig.WebPort)

	http.ListenAndServe(W.webConfig.WebAddress+":"+W.webConfig.WebPort, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"*"}),
	)(handler))
}

//New creates a new handler
func (W *APIOne) New() http.Handler {

	log.Infoln("Opening Database")
	//STAGE 2 Open Database
	DBClient := storage.Client{}
	//Include datbase separators
	DB := DBClient.OpenBoltDb(constants.ContentDir, W.webConfig.DatabaseName)
	DB.Seed()

	app, err := webapi.NewApp(W.webConfig, DB)

	if err != nil {
		log.Fatalln("Error creating API ")
		return nil
	}

	api := app.Mux.PathPrefix("/launchapi/api/v1").Subrouter()

	api.HandleFunc("/liveness", app.Liveness)
	api.HandleFunc("/about", app.About)

	//Rocket Section
	api.Handle("/rocket/getAll", app.AuthMiddleware(http.HandlerFunc(app.RocketGetALL)))

	//Orbit Section
	api.Handle("/orbit/getAll", app.AuthMiddleware(http.HandlerFunc(app.OrbitGetALL)))

	//SpacePorts Section
	api.Handle("/spaceport/getAll", app.AuthMiddleware(http.HandlerFunc(app.SpacePortGetALL)))

	//Engines Section
	api.Handle("/engine/getAll", app.AuthMiddleware(http.HandlerFunc(app.EngineGetALL)))

	api.HandleFunc("/login", app.V1Login)
	api.HandleFunc("/logout", app.V1Logout)
	api.HandleFunc("/register", app.V1CreateUser)

	return &app

}
