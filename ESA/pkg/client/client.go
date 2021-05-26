package client

import (
	"context"
	"encoding/json"
	"epyphite/space/v1/ESA/pkg/models"
	"epyphite/space/v1/ESA/pkg/models/modules"
	"errors"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
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
	log.Debugf("Making Rquest to ->  %s \n", req.URL)
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
		log.Debugln(err)
		return err
	}

	return nil
}

func (c *RestClient) GetDiscusObjects(ctx context.Context, options *modules.DiscosRequest) ([]*modules.DiscosResponse, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, options.Prefix)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error processing URL %s, \n Error: %s", url, err)
	}
	req = req.WithContext(ctx)
	res := []*modules.DiscosResponse{}
	err = c.sendRequest(req, &res)
	if err != nil {
		log.Errorln("Error on NeoW  Module: ", err)
	}
	return res, err
}
