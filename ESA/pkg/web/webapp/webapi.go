package webapp

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"

	"github.com/gorilla/sessions"

	explorer "epyphite/space/v1/ESA"
	models "epyphite/space/v1/ESA/pkg/models"
	"epyphite/space/v1/ESA/pkg/models/modules"
	"epyphite/space/v1/ESA/pkg/storage"
	c1 "epyphite/space/v1/ESA/pkg/web/constants"
)

//JResponseDiscuss create a trscture to respond json
type JResponseDiscuss struct {
	ResponseCode string
	Message      string
	ResponseData []*modules.DiscosResponse
}

//JResponse create a trscture to respond json
type JResponse struct {
	ResponseCode string
	Message      string
	ResponseData string
}

//MainWebAPI PHASE
type MainWebAPI struct {
	Mux     *mux.Router
	Config  models.Config
	storage *storage.Client
	Store   *sessions.CookieStore
}

//GetFileContentType will get the mime type of the file by reading its first 512 bytes (according to the standard)
func GetFileContentType(buffer []byte) (string, error) {
	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

//NewApp create a new object for the App.
func NewApp(config models.Config, db *storage.Client) (MainWebAPI, error) {
	var err error
	var wapp MainWebAPI

	mux := mux.NewRouter().StrictSlash(true)
	wapp.storage = db

	wapp.Mux = mux
	wapp.Config = config

	if err != nil {
		log.Println(err)
	}
	return wapp, err
}

func (a *MainWebAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Mux.ServeHTTP(w, r)
}

func getSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := c1.Store.Get(r, "esa-session")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return nil
	}
	return session
}

//Liveness just keeps the connection alive
func (a *MainWebAPI) Liveness(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var response JResponse

	response.Message = "Process Alive"
	response.ResponseCode = "200"
	response.ResponseData = ""
	js, err := json.Marshal(response)
	if err != nil {
		log.Println()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

//About just keeps the connection alive
func (a *MainWebAPI) About(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var response JResponse

	response.Message = "Version 1"
	response.ResponseCode = "200"
	response.ResponseData = c1.AboutText
	js, err := json.Marshal(response)
	if err != nil {
		log.Println()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

//GetAllTLECollection just keeps the connection alive
func (a *MainWebAPI) GetDiscusALL(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	lteRet, err := explorer.GetDISCUSALL(a.Config)
	var response JResponseDiscuss
	response.Message = "Version 1"
	response.ResponseCode = "200"
	response.ResponseData = lteRet
	js, err := json.Marshal(response)
	if err != nil {
		log.Println()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}
