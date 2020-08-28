package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	models "github.com/epyphite/space/NASA/pkg/models"
	modules "github.com/epyphite/space/NASA/pkg/models/modules"
)

//RestClient is the rest client used for connecting to a rest API
type RestClient struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

//NewClient returns a new client with the required information.
func NewClient(config models.Config) *RestClient {
	return &RestClient{
		APIKey:  config.APIKey,
		BaseURL: config.BaseURL,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *RestClient) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	fmt.Printf("Making Rquest to ->  %s \n", req.URL)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("SendRequest -> on making request to %s - %s ", req.URL, err)
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}
		return fmt.Errorf("unknown error, status code: %d, url: %s", res.StatusCode, req.URL)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

//EONET

//GetEonetEvent will retrieve data events from the Eonet Network
func (c *RestClient) GetEonetEvent(ctx context.Context, options *modules.EonetRequest) (*modules.EonetEventResponse, error) {
	url := fmt.Sprintf("%s/%s?days=%d&status=%s", c.BaseURL, options.Prefix, options.Days, options.Status)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error processing URL %s, \n Error: %s", url, err)
	}

	req = req.WithContext(ctx)
	res := modules.EonetEventResponse{}
	err = c.sendRequest(req, &res)
	if err != nil {
		log.Println("Error on NeoW  Module: ", err)
	}
	return &res, err
}

//GetEonetCategory will retrieve EONET Categories
func (c *RestClient) GetEonetCategory(ctx context.Context, options *modules.EonetRequest) (*modules.EonetCategoryResponse, error) {
	url := fmt.Sprintf("%s/%s?days=%d&status=%s", c.BaseURL, options.Prefix, options.Days, options.Status)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error processing URL %s, \n Error: %s", url, err)
	}

	req = req.WithContext(ctx)
	res := modules.EonetCategoryResponse{}
	err = c.sendRequest(req, &res)
	if err != nil {
		log.Println("Error on NeoW  Module: ", err)
	}
	return &res, err
}

//END EONET

//GetNeoBrowse will bring the Near Object Database
func (c *RestClient) GetNeoBrowse(ctx context.Context, options *modules.NeoWBrowseRequest) (*modules.NeoWBroseResponse, error) {
	url := fmt.Sprintf("%s/%s?api_key=%s", c.BaseURL, options.Prefix, c.APIKey)

	if options.Page != 0 {
		url = fmt.Sprintf("%s/%s/?page=%d&api_key=%s", c.BaseURL, options.Prefix, options.Page, c.APIKey)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error processing URL %s, \n Error: %s", url, err)
	}

	req = req.WithContext(ctx)
	res := modules.NeoWBroseResponse{}
	err = c.sendRequest(req, &res)
	if err != nil {
		log.Println("Error on NeoW  Module: ", err)
	}
	return &res, err
}

func (c *RestClient) GetNeoLookUp(ctx context.Context, options *modules.NeoLookUpRequest) (*modules.NearEarthObject, error) {
	url := fmt.Sprintf("%s/%s/%s?api_key=%s", c.BaseURL, options.Prefix, options.AsteroidID, c.APIKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error processing URL %s, \n Error: %s", url, err)
	}

	req = req.WithContext(ctx)
	res := modules.NearEarthObject{}
	err = c.sendRequest(req, &res)
	if err != nil {
		log.Println("Error on NeoW  Module: ", err)
	}
	return &res, err
}

func (c *RestClient) GetNeoWFeed(ctx context.Context, options *modules.NeoWFeedRequest) (*modules.NeoWFeedResponse, error) {

	if options.StartDate == "" {
		return nil, fmt.Errorf("Error in GetNeoW you need to specify a start date")
	}
	url := fmt.Sprintf("%s/%s?start_date=%s&api_key=%s", c.BaseURL, options.Prefix, options.StartDate, c.APIKey)

	if options.EndDate != "" {
		url = fmt.Sprintf("%s/%s?start_date=%s&end_date=%s&api_key=%s", c.BaseURL, options.Prefix, options.StartDate, options.EndDate, c.APIKey)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error processing URL %s, \n Error: %s", url, err)
	}

	req = req.WithContext(ctx)
	res := modules.NeoWFeedResponse{}
	err = c.sendRequest(req, &res)

	if err != nil {
		log.Println("Error on NeoW  Module: ", err)
	}
	return &res, err
}

//GetAPOD will get a Astronomical pciture of the day.
func (c *RestClient) GetAPOD(ctx context.Context, options *modules.ApodRequest) (*modules.ApodResponse, error) {

	if options.Date == "today" {
		input := time.Now()
		layout := "2006-01-02"
		t := input.Format(layout)
		options.Date = t
	}
	url := fmt.Sprintf("%s/%s?date=%s&hd=%t&api_key=%s", c.BaseURL, options.Prefix, options.Date, options.HD, c.APIKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error processing URL %s, \n Error: %s", url, err)
	}

	req = req.WithContext(ctx)
	res := modules.ApodResponse{}
	err = c.sendRequest(req, &res)

	if err != nil {
		fmt.Println("Error on APOD Module: ", err)
	}
	return &res, err
}

//GetTLECollection will save the required TLE pages
func (c *RestClient) GetTLECollection(ctx context.Context, options *modules.TLECollectionRequest) (*modules.TLECollectionResponse, error) {

	var url string
	var page int
	var err error

	page = options.Page
	if page > 0 {
		url = fmt.Sprintf("%s/%s?page=%d", c.BaseURL, options.Prefix, page)
	} else {
		url = fmt.Sprintf("%s/%s", c.BaseURL, options.Prefix)
	}

	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error processing URL %s, \n Error: %s", url, err)
	}

	req = req.WithContext(ctx)
	res := modules.TLECollectionResponse{}
	err = c.sendRequest(req, &res)

	if err != nil {
		fmt.Println("Error on TLE Module: ", err)
	}
	return &res, err

}
