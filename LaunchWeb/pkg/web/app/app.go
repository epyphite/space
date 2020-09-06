package MainWebApp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

	"github.com/epyphite/space/LaunchWeb/pkg/models"
	c1 "github.com/epyphite/space/LaunchWeb/pkg/web/constants"
	"github.com/mailgun/mailgun-go"

	"github.com/gorilla/sessions"
)

//Claims Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//Credentials Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

//JResponse create a trscture to respond json
type JResponse struct {
	Status  string
	Message string
}

//Requirement
type Inquiry struct {
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Select_field string `json:"select"`
	Terms        string `json:"terms"`
}

type Message struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Terms   string `json:"terms"`
}

//MainWebApp PHASE
type MainWebApp struct {
	Mux    *mux.Router
	Log    *log.Logger
	Config models.Config
}

//NewApp creates a new instance
func NewApp(config models.Config) (MainWebApp, error) {
	var err error
	var wapp MainWebApp

	mux := mux.NewRouter().StrictSlash(true)

	log.SetFormatter(&log.JSONFormatter{})
	wapp.Mux = mux
	wapp.Config = config
	wapp.Log = &log.Logger{}

	if err != nil {
		log.Println(err)
	}

	return wapp, err
}

func getSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := c1.Store.Get(r, "iochu-session")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return nil
	}
	return session
}

//ServeIndex will serve web pages
func (a *MainWebApp) ServeIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, a.Config.ContentDir+"/index.html")

}

func (a *MainWebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Mux.ServeHTTP(w, r)
}

//SendMessage send messages on MG
func (a *MainWebApp) SendMessage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/sendMessage" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var domain string
	var apiKey string

	log.Println("SendSimpleMessage")

	domain = os.Getenv("MAILGUN_DOMAIN")
	apiKey = os.Getenv("MAILGUN_APIKEY")

	var message Message
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&message)
	var body string
	body = message.Name + "\n" + message.Email + "\n" + message.Message
	mg := mailgun.NewMailgun(domain, apiKey)

	m := mg.NewMessage(
		"Excited User <mailgun@epyphite.com>",
		"Message from Landing Page",
		body,
		"contact@epyphite.com",
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, _, err = mg.Send(ctx, m)

	log.Println("Sended response")
	response := JResponse{"successful", "Login Successful"}
	if err != nil {
		response.Message = err.Error()
		response.Status = "Error"
		log.Println("Error Sending Message ", err)
	}
	js, err := json.Marshal(response)

	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

//SendInquiry for send inquiry
func (a *MainWebApp) SendInquiry(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/sendInquiry" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	var domain string
	var apiKey string
	log.Println("sendInquiry ")

	domain = os.Getenv("MAILGUN_DOMAIN")
	apiKey = os.Getenv("MAILGUN_APIKEY")

	var inquiry Inquiry
	fmt.Println(r.Body)

	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&inquiry)
	fmt.Println(r.Body)
	if err != nil {
		log.Println(err)
	}
	var body string

	body = inquiry.Name + "\n" + inquiry.Email + "\n" + inquiry.Phone + "\n" + inquiry.Select_field

	mg := mailgun.NewMailgun(domain, apiKey)

	m := mg.NewMessage(
		"Excited User <mailgun@epyphite.com>",
		"Inquiry from Landing Page",
		body,
		"contact@epyphite.com",
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, _, err = mg.Send(ctx, m)

	log.Println("Sended response")
	response := JResponse{"successful", "Login Successful"}
	if err != nil {
		response.Message = err.Error()
		response.Status = "Error"
		log.Println("Error Sending Message ", err)
	}
	js, err := json.Marshal(response)

	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}
