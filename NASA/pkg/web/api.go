package webapi

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"

	constants "github.com/epyphite/space/NASA/pkg/constants"
	models "github.com/epyphite/space/NASA/pkg/models"
	"github.com/epyphite/space/NASA/pkg/storage"
	webapi "github.com/epyphite/space/NASA/pkg/web/webapp"
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
	DB := DBClient.OpenBoltDb("./data", W.webConfig.DatabaseName)
	DB.Seed()

	app, err := webapi.NewApp(W.webConfig, DB)

	if err != nil {
		log.Fatalln("Error creating API ")
		return nil
	}

	api := app.Mux.PathPrefix("/nasa/api/v1").Subrouter()

	api.HandleFunc("/liveness", app.Liveness)
	api.HandleFunc("/about", app.About)
	api.HandleFunc("/tle/getall", app.GetAllTLECollection)
	api.HandleFunc("/apod/get", app.GetAPOD)
	api.HandleFunc("/NEO/getall", app.GetNeoAll)

	return &app

}
