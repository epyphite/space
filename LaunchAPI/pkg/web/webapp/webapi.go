package webapp

import (
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"

	"github.com/gorilla/sessions"

	models "github.com/epyphite/space/LaunchAPI/pkg/models"
	"github.com/epyphite/space/LaunchAPI/pkg/storage"
	c1 "github.com/epyphite/space/LaunchAPI/pkg/web/constants"
)

//JResponse create a trscture to respond json
type JResponse struct {
	ResponseCode string
	Message      string
	ResponseData []byte
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

func (a *MainWebAPI) getSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := c1.Store.Get(r, "spaceLaunch-session")
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
	response.ResponseData = []byte("")
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
	response.ResponseData = []byte(c1.AboutText)
	js, err := json.Marshal(response)
	if err != nil {
		log.Println()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

//RocketGetALL will get all Rocket data loaded
func (a *MainWebAPI) RocketGetALL(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	response, err := a.storage.RocketGetAll()

	js, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		log.Println()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

//SpacePortGetALL will get all Rocket data loaded
func (a *MainWebAPI) SpacePortGetALL(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	response, err := a.storage.SpacePortGetAll()

	js, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		log.Println()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

//OrbitGetALL will get all Rocket data loaded
func (a *MainWebAPI) OrbitGetALL(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	response, err := a.storage.OrbitGetAll()

	js, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		log.Println()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

//EngineGetALL will get all Rocket data loaded
func (a *MainWebAPI) EngineGetALL(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	response, err := a.storage.EngineGetAll()

	js, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		log.Println()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

/// System Calls

//V1Login main login function to keep also store
func (a *MainWebAPI) V1Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting response before options")

	setupResponse(&w, r)

	log.Println("Getting response before options")
	if r.Method == "OPTIONS" {
		return
	}
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var err error

	session, err := a.Store.Get(r, "qplace-go-session")

	var _user models.JSONLogin
	var user models.User
	err = json.NewDecoder(r.Body).Decode(&_user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("User ", _user.Email)

	user.Email = _user.Email
	user.Password = []byte(_user.Password)

	user, auth, err := a.storage.CheckUser(user)

	if err != nil {
		/*
			var response JResponse
			response.ResponseCode = "201"
			response.Message = "incorrect Username or Password "
			response.ResponseData = []byte("")
			jresponse, err := json.Marshal(response)
		*/
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "Application/json")
		//w.Write(jresponse)
		http.Error(w, "Not authorized", 401)

		return
	}
	if auth {
		if user.Banned == true {
			w.Header().Set("Content-Type", "Application/json")
			//w.Write(jresponse)
			http.Error(w, "Not authorized", 401)

			return
		}

		if user.Approved == false {
			w.Header().Set("Content-Type", "Application/json")
			//w.Write(jresponse)
			http.Error(w, "Not authorized", 401)

			return
		}
		user.Password = []byte("") // We empty the password
		var usersResponse []models.User
		var response models.JSONResponseUsers

		usersResponse = append(usersResponse, user)
		response.ResponseCode = "200"
		response.Message = "logged in Succesfully"
		response.ResponseData = usersResponse

		jresponse, err := json.Marshal(response)

		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "Application/json")
		session.Options.Path = "/"
		session.Options.MaxAge = 3600
		session.Options.HttpOnly = true
		session.Values["user"] = user.UserID
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(jresponse)
		return
	} else {

		/*
			var response JResponse
			response.ResponseCode = "201"
			response.Message = "incorrect Username or Password "
			response.ResponseData = []byte("")
			jresponse, err := json.Marshal(response)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "Application/json")
			w.Write(jresponse)
		*/
		http.Error(w, err.Error(), http.StatusForbidden)

		return
	}

}

//V1Logout destro session
func (a *MainWebAPI) V1Logout(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	session := a.getSession(w, r)
	w.Header().Set("Content-Type", "Application/json")
	session.Options.Path = "/"
	session.Options.MaxAge = -1 //We remove the session completely
	session.Options.HttpOnly = true
	a.Store.Save(r, w, session)
	w.Header().Set("Content-Type", "Application/json")
	http.Redirect(w, r, "/", 301)

}

//CheckSession validates that user has a session active
func (a *MainWebAPI) CheckSession(w http.ResponseWriter, r *http.Request) bool {
	setupResponse(&w, r)

	session := a.getSession(w, r)
	// MOCK function we should add server status , this is a TEST WIP TODO session

	userID, found := session.Values["user"]
	if !found {
		fmt.Println("No user_id found in session")
		return false
	}

	str := fmt.Sprintf("%v", userID)
	user, err := a.storage.CheckUserByID(str)

	if err != nil {
		log.Println("Session Failed to renew or Expired")
		http.Error(w, "unauthorised", http.StatusUnauthorized)
		return false
	}
	if user.Email == "" {
		log.Println("Session Failed to renew or Expired")
		http.Error(w, "unauthorised", http.StatusUnauthorized)
		return false
	}

	a.Store.MaxAge(3600) // renew session 1 Minute
	a.Store.Save(r, w, session)
	return true
}

//V1CreateUser destro session
func (a *MainWebAPI) V1CreateUser(w http.ResponseWriter, r *http.Request) {
	var err error
	setupResponse(&w, r)
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var response JResponse
	log.Printf(r.Header.Get("AdminToken"))
	log.Println(r.Header)

	var _user models.JSONCreateUser

	err = json.NewDecoder(r.Body).Decode(&_user)

	var tempUser models.User
	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		response.ResponseCode = "401"
		response.Message = "Error Creating user "
		response.ResponseData = []byte("")
	}
	tempUser.Username = _user.UserName
	tempUser.Email = _user.Email
	tempUser.Password = []byte(_user.Password) //Default Password CHANGE IN PROD
	tempUser.UserID = u2.String()
	tempUser.Token = ""
	tempUser.Approved = false
	tempUser.Banned = false
	tempUser.Role = "Admin"
	err = a.storage.UserAdd(tempUser)
	if err != nil {
		response.ResponseCode = "401"
		response.Message = "Error Creating user "
		response.ResponseData = []byte("")
	}
	response.ResponseCode = "201"
	response.Message = "User Created"
	response.ResponseData = []byte("")

	/*
		_token := r.Header.Get("AdminToken")

		if _token != "" {
			token, err := a.storage.CheckToken(_token)
			fmt.Println(token != (models.Token{}))
			fmt.Println(err)

			if err != nil {
				fmt.Println("Error ", err)
			}
			if token != (models.Token{}) {
				if token.IsAdmin {


				} else {
					response.ResponseCode = "401"
					response.Message = "Token is not admin"
					response.ResponseData = []byte("")
				}
			} else {
				response.ResponseCode = "401"
				response.Message = "Error Creating user, token not found "
				response.ResponseData = []byte("")
			}

		} else {
			response.ResponseCode = "201"
			response.Message = "Error Creating user "
			response.ResponseData = []byte("")
		}

	*/

	jresponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	if response.ResponseCode == "401" {
		http.Error(w, "Not authorized", 401)

	} else {
		w.Write(jresponse)
	}

	return

}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	log.Println("setting up")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "admintoken, Content,Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
