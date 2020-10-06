package utils

import (
	"encoding/json"
	"net/url"
	"os"

	"epyphite/space/v1/NASA/pkg/models"
)

//GetVarURL will return the value of a URL Parameter
func GetVarURL(uri string, parameter string) string {

	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	return (m[parameter][0])
}

//LoadConfiguration returns the read Configuration and error while reading.
func LoadConfiguration(file string) (models.Config, error) {
	var config models.Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
}
