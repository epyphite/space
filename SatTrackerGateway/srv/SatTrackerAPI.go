package srv

import (
	"bytes"
	"encoding/json"
	"epyphite/space/v1/SatTrackerGateway/pkg/constants"
	"epyphite/space/v1/SatTrackerGateway/pkg/models"
	"epyphite/space/v1/SatTrackerGateway/pkg/storage"
	"os"

	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/publicsuffix"

	"net/http"
	"net/http/cookiejar"
)

type SatTrackAPI struct {
	Client http.Client
}

//Login Function
func (ST *SatTrackAPI) Login(identity models.Identity) ([]byte, error) {
	requestBody, err := json.Marshal(identity)
	if err != nil {
		log.Fatalln(err)
	}
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		log.Fatal(err)
	}
	ST.Client = http.Client{Jar: jar}
	log.Println(constants.BASEURI + constants.LOGINURI)
	resp, err := ST.Client.Post(constants.BASEURI+constants.LOGINURI, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode != 200 {
		log.Println("Status Code ", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body, err
}

//MakeRequestBoxScore to the StarTracker services
func (ST *SatTrackAPI) MakeRequestBoxScore() ([]models.SatTrackBasic, error) {

	var err error
	log.Println(constants.BASEURI + constants.BoxScore)

	resp, err := ST.Client.Get(constants.BASEURI + constants.BoxScore)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	var result []models.SatTrackBasic

	client, err := storage.NewClient(os.Getenv("EPY_MONGOURI"))

	json.NewDecoder(resp.Body).Decode(&result)
	for _, r := range result {
		insertedID, err := client.SaveSatObjectBasic(r)
		if err != nil {
			log.Println(err)
		}
		log.Println("Saved ID ", insertedID)
	}
	return result, err
}

//MakeRequestLeoCurrent to the StarTracker services
func (ST *SatTrackAPI) MakeRequestLeoCurrent() ([]models.SatTrackStandard, error) {

	var err error
	log.Println(constants.BASEURI + constants.CurrentLEO)

	resp, err := ST.Client.Get(constants.BASEURI + constants.CurrentLEO)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	var result []models.SatTrackStandard

	client, err := storage.NewClient(os.Getenv("EPY_MONGOURI"))

	json.NewDecoder(resp.Body).Decode(&result)
	for _, r := range result {
		insertedID, err := client.SaveSatObject(r)
		if err != nil {
			log.Println(err)
		}
		log.Println("Saved ID ", insertedID)
	}
	return result, err
}
