//+build integration

package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	models "github.com/epyphite/space/NASA/pkg/models"
	modules "github.com/epyphite/space/NASA/pkg/models/modules"
	"github.com/stretchr/testify/assert"
)

func TestGetTLECollection(t *testing.T) {
	config := models.Config{APIKey: os.Getenv("NASA_KEY")}
	config.BaseURL = "https://data.ivanstanojevic.me"

	TLEOptions := modules.TLECollectionRequest{Prefix: "api/tle"}

	c := NewClient(config)
	ctx := context.Background()

	res, err := c.GetTLECollection(ctx, &TLEOptions)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	file, _ := json.MarshalIndent(res, "", " ")

	_ = ioutil.WriteFile("test_TestGetTLECollection.json", file, 0644)
}

func TestGetApod(t *testing.T) {

	config := models.Config{APIKey: os.Getenv("NASA_KEY")}
	config.BaseURL = "https://api.nasa.gov"

	apodOptions := modules.ApodRequest{Date: "today", HD: true, Prefix: "planetary/apod"}
	c := NewClient(config)
	ctx := context.Background()
	res, err := c.GetAPOD(ctx, &apodOptions)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	assert.NotNil(t, res.URL, "expecting value on url")
}

func TestGetNeoWFeed(t *testing.T) {

	config := models.Config{APIKey: os.Getenv("NASA_KEY")}
	config.BaseURL = "https://api.nasa.gov"

	neoOptions := modules.NeoWFeedRequest{StartDate: "2020-08-11", EndDate: "", Prefix: "neo/rest/v1/feed"}
	c := NewClient(config)
	ctx := context.Background()
	res, err := c.GetNeoWFeed(ctx, &neoOptions)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	file, _ := json.MarshalIndent(res, "", " ")

	_ = ioutil.WriteFile("test_TestGetNeoWFeed.json", file, 0644)
}

func TestGetNeoWLookup(t *testing.T) {
	config := models.Config{APIKey: os.Getenv("NASA_KEY")}
	config.BaseURL = "https://api.nasa.gov"

	neoOptions := modules.NeoLookUpRequest{AsteroidID: "3542519", Prefix: "neo/rest/v1/neo"}
	c := NewClient(config)
	ctx := context.Background()
	res, err := c.GetNeoLookUp(ctx, &neoOptions)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	file, _ := json.MarshalIndent(res, "", " ")

	_ = ioutil.WriteFile("test_TestGetNeoWLookup.json", file, 0644)
}

func TestGetNeoWBrowse(t *testing.T) {
	config := models.Config{APIKey: os.Getenv("NASA_KEY")}
	config.BaseURL = "https://api.nasa.gov"

	neoOptions := modules.NeoWBrowseRequest{Prefix: "neo/rest/v1/neo/browse"}
	c := NewClient(config)
	ctx := context.Background()
	res, err := c.GetNeoBrowse(ctx, &neoOptions)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	file, _ := json.MarshalIndent(res, "", " ")

	_ = ioutil.WriteFile("test_TestGetNeoWBrowse.json", file, 0644)
}

func TestGetEonetEvent(t *testing.T) {
	config := models.Config{APIKey: os.Getenv("NASA_KEY")}
	config.BaseURL = "https://eonet.sci.gsfc.nasa.gov"

	eonetOptions := modules.EonetRequest{Prefix: "/api/v2.1/events", Days: 20, Status: "open"}
	c := NewClient(config)
	ctx := context.Background()
	res, err := c.GetEonetEvent(ctx, &eonetOptions)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	file, _ := json.MarshalIndent(res, "", " ")

	_ = ioutil.WriteFile("test_TestGetEonetEvent.json", file, 0644)
}
func TestGetEonetCategory(t *testing.T) {
	config := models.Config{APIKey: os.Getenv("NASA_KEY")}
	config.BaseURL = "https://eonet.sci.gsfc.nasa.gov"

	eonetOptions := modules.EonetRequest{Prefix: "/api/v2.1/categories", Days: 20, Status: "open"}

	c := NewClient(config)
	ctx := context.Background()

	res, err := c.GetEonetCategory(ctx, &eonetOptions)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	file, _ := json.MarshalIndent(res, "", " ")

	_ = ioutil.WriteFile("test_TestGetEonetCategory.json", file, 0644)
}
