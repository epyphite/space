package nasaexplorer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	client "epyphite/space/v1/NASA/pkg/client"
	models "epyphite/space/v1/NASA/pkg/models"
	modules "epyphite/space/v1/NASA/pkg/models/modules"
	parser "epyphite/space/v1/NASA/pkg/parser"
	"epyphite/space/v1/NASA/pkg/srv"

	"epyphite/space/v1/NASA/pkg/utils"
)

//GetLatestApod Get Picture of the day from nasa, API key is required
func GetLatestApod(configuration models.Config) (*modules.ApodResponse, error) {

	var config models.Config
	config = configuration
	if config == (models.Config{}) {
		config = models.Config{APIKey: os.Getenv("NASA_KEY")}
	}

	if config.BaseURL == "" {
		config.BaseURL = "https://api.nasa.gov"
	}
	apodOptions := modules.ApodRequest{Date: "today", HD: true, Prefix: "planetary/apod"}
	c := client.NewClient(config)

	ctx := context.Background()
	res, err := c.GetAPOD(ctx, &apodOptions)
	return res, err
}

//GetEonetLatestEvent retrieves events from the Eonet API with the last 20 days
// default values:
// 20 days
// status open
func GetEonetLatestEvent(configuration models.Config) (*modules.EonetEventResponse, error) {
	var config models.Config
	config = configuration
	if config == (models.Config{}) {
		config = models.Config{APIKey: os.Getenv("NASA_KEY")}
	}

	if config.BaseURL == "" {
		config.BaseURL = "https://eonet.sci.gsfc.nasa.gov"
	}
	eonetOptions := modules.EonetRequest{Prefix: "/api/v2.1/events", Days: 20, Status: "open"}
	c := client.NewClient(config)
	ctx := context.Background()
	res, err := c.GetEonetEvent(ctx, &eonetOptions)
	return res, err
}

//GetEonetListCategories retrieves categories
func GetEonetListCategories(configuration models.Config) (*modules.EonetCategoryResponse, error) {

	var eonetOptions modules.EonetRequest
	var config models.Config
	config = configuration
	if config == (models.Config{}) {
		config = models.Config{APIKey: os.Getenv("NASA_KEY")}
	}

	if config.BaseURL == "" {
		config.BaseURL = "https://eonet.sci.gsfc.nasa.gov"
	}

	eonetOptions = modules.EonetRequest{Prefix: "/api/v2.1/categories", Days: 20, Status: "open"}

	c := client.NewClient(config)
	ctx := context.Background()

	res, err := c.GetEonetCategory(ctx, &eonetOptions)
	return res, err
}

//GetNeoAll will bring all the Near Earth Objects up to "page",
// Configure also save on error to allow big queries being saved
func GetNeoAll(configuration models.Config) ([]*modules.NeoWBroseResponse, error) {
	var config models.Config
	config = configuration
	var neoResponse []*modules.NeoWBroseResponse
	if config == (models.Config{}) {
		config = models.Config{APIKey: os.Getenv("NASA_KEY")}
	}

	if config.BaseURL == "" {
		config.BaseURL = "https://api.nasa.gov"
	}

	c := client.NewClient(config)
	ctx := context.Background()
	neoOptions := modules.NeoWBrowseRequest{Prefix: "neo/rest/v1/neo/browse"}

	res, err := c.GetNeoBrowse(ctx, &neoOptions)

	if err != nil {
		return nil, err
	}
	neoResponse = append(neoResponse, res)

	for res.Links.Next != "" {
		page, _ := strconv.Atoi(utils.GetVarURL(res.Links.Next, "page"))
		if page >= config.MaxPages {
			break
		}
		neoOptions := modules.NeoWBrowseRequest{Prefix: "neo/rest/v1/neo/browse", Page: page}
		res, err = c.GetNeoBrowse(ctx, &neoOptions)
		if err != nil {
			log.Println("Error retrieving the next page")
			if config.SaveOnError == true {
				break
			}
			return nil, err
		}
		neoResponse = append(neoResponse, res)
	}
	return neoResponse, err

}

//GetAllTLECollection will save the required TLE pages
func GetAllTLECollection(configuration models.Config) ([]*modules.TLECollectionResponse, error) {
	var config models.Config
	config = configuration
	var lteCollection []*modules.TLECollectionResponse
	if config == (models.Config{}) {
		config = models.Config{APIKey: os.Getenv("NASA_KEY")}
	}

	if config.BaseURL == "" {
		config.BaseURL = "https://data.ivanstanojevic.me"
	}

	c := client.NewClient(config)
	ctx := context.Background()

	tleOptions := modules.TLECollectionRequest{Prefix: "api/tle"}

	res, err := c.GetTLECollection(ctx, &tleOptions)

	lteCollection = append(lteCollection, res)
	for res.View.Next != "" {
		page, _ := strconv.Atoi(utils.GetVarURL(res.View.Next, "page"))
		if page >= config.MaxPages {
			break
		}
		tleOptions := modules.TLECollectionRequest{Prefix: "api/tle", Page: page}
		res, err = c.GetTLECollection(ctx, &tleOptions)

		if err != nil {
			log.Println("Error retrieving the next page")
			if config.SaveOnError == true {
				break
			}
			return nil, err
		}
		lteCollection = append(lteCollection, res)
		fmt.Println(len(lteCollection))
	}
	return lteCollection, err

}

//GetTLEMember will save the required TLE pages
func GetTLEMember(configuration models.Config, satID int) (*modules.TLEMember, error) {
	var config models.Config
	config = configuration
	if config == (models.Config{}) {
		config = models.Config{APIKey: os.Getenv("NASA_KEY")}
	}

	if config.BaseURL == "" {
		config.BaseURL = "https://data.ivanstanojevic.me"
	}

	c := client.NewClient(config)
	ctx := context.Background()

	tleOptions := modules.TLERecordRequest{Prefix: "api/tle", ID: satID}

	res, err := c.GetTLEMember(ctx, &tleOptions)

	return res, err

}

//GetTLEMemberDetails will save the required TLE pages
func GetTLEMemberDetails(configuration models.Config, satID int) (*modules.TLEMemberFormat, error) {
	var config models.Config
	config = configuration
	if config == (models.Config{}) {
		config = models.Config{APIKey: os.Getenv("NASA_KEY")}
	}

	if config.BaseURL == "" {
		config.BaseURL = "https://data.ivanstanojevic.me"
	}

	c := client.NewClient(config)
	ctx := context.Background()

	tleOptions := modules.TLERecordRequest{Prefix: "api/tle", ID: satID}

	res, err := c.GetTLEMember(ctx, &tleOptions)

	resMember := parser.ParseTLEMember(res)

	// Added to Orbital Element
	var OE srv.OrbitalElement

	OE.Anomaly = resMember.Line2.MeanAnomaly
	OE.Inclination = resMember.Line2.OrbitInclination
	OE.Ascension = resMember.Line2.RAAN
	OE.Eccentricity = resMember.Line2.Eccentricity
	OE.Motion = resMember.Line2.MeanMotion
	OE.Perigee = resMember.Line2.ArgumentPerigee
	OE.CalculateSemiMajorAxis()
	OE.CalculateTrueAnomaly()
	var file []byte
	file, _ = json.MarshalIndent(OE, "", " ")
	_ = ioutil.WriteFile("OE.json", file, 0644)
	return resMember, err
}
