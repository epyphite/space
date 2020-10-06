package web

import (
	"log"
	"net/http"

	models "epyphite/space/v1/LaunchWeb/pkg/models"
	"github.com/gorilla/handlers"

	constants "epyphite/space/v1/LaunchWeb/pkg/constants"
	webapp "epyphite/space/v1/LaunchWeb/pkg/web/app"
	"epyphite/space/v1/LaunchWeb/pkg/web/ui"
)

type WebOne struct {
	webconfig models.Config
}

//NewWebAgent // creates a mew instace \of web one
func NewWebAgent(config models.Config) (WebOne, error) {
	var webOne WebOne
	webOne.webconfig = config
	return webOne, nil
}

//StartServer Starts the server using the variable sip and port, creates anew instance.
func (W *WebOne) StartServer() {
	log.Println("Version : " + constants.BuildVersion)
	log.Println("Build Time : " + constants.BuildTime)
	handler := W.New()
	http.ListenAndServe(W.webconfig.WebAddress+":"+W.webconfig.WebPort, handlers.CORS(handlers.AllowedOrigins([]string{"*"}), handlers.AllowedMethods([]string{"*"}))(handler))
}

//New creates a new handler
func (W *WebOne) New() http.Handler {

	app, err := webapp.NewApp(W.webconfig)

	if err != nil {
		log.Fatalln("Error creating webapp ")
		return nil
	}
	api := app.Mux.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/sendMessage", app.SendMessage)
	api.HandleFunc("/sendInquiry", app.SendInquiry)

	UIHandler := ui.UI()
	app.Mux.PathPrefix("/js").Handler(http.StripPrefix("/", UIHandler))
	app.Mux.PathPrefix("/css").Handler(http.StripPrefix("/", UIHandler))
	app.Mux.PathPrefix("/img").Handler(http.StripPrefix("/", UIHandler))
	app.Mux.PathPrefix("/data").Handler(http.StripPrefix("/", UIHandler))

	app.Mux.PathPrefix("/favicon.ico").Handler(http.StripPrefix("/", UIHandler))
	app.Mux.PathPrefix("/").Handler(UIHandler)

	// Stats Call , external Func
	return &app
}
