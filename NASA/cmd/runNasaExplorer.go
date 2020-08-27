package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	explorer "github.com/epyphite/space/NASA"
	"github.com/epyphite/space/NASA/pkg/models"
)

func main() {
	options := models.Config{}
	options.APIKey = os.Getenv("NASA_KEY")

	apodRet, err := explorer.GetLatestApod(options)
	if err != nil {
		fmt.Println(err)
	}
	file, _ := json.MarshalIndent(apodRet, "", " ")
	_ = ioutil.WriteFile("Apod.json", file, 0644)

	eonetRet, err := explorer.GetEonetLatestEvent(options)
	if err != nil {
		fmt.Println(err)
	}
	file, _ = json.MarshalIndent(eonetRet, "", " ")
	_ = ioutil.WriteFile("Eonet.json", file, 0644)

	options.SaveOnError = true
	options.MaxPages = 10
	neoRet, err := explorer.GetNeoAll(options)
	if err != nil {
		fmt.Println(err)
	}
	file, _ = json.MarshalIndent(neoRet, "", " ")
	_ = ioutil.WriteFile("Neo.json", file, 0644)

}
