package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	explorer "github.com/epyphite/space/NASA"
	"github.com/epyphite/space/NASA/pkg/models"
	"github.com/epyphite/space/NASA/pkg/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nasaExplorer",
	Short: "nasaExplorer",
	Long:  ``,
	RunE:  nasaExplorer,
}
var cfgFile string
var services []string
var tleSatelite int

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Define a configuration file location")
	rootCmd.PersistentFlags().StringArrayVar(&services, "services", nil, "Define the services you wish to run")
	rootCmd.PersistentFlags().IntVar(&tleSatelite, "satid", -1, "Define a tle Satelite")

}

//Execute will run the desire module command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func nasaExplorer(cmd *cobra.Command, args []string) error {

	execName, _ := os.Executable()
	execName = filepath.Base(execName)
	options := models.Config{}
	var err error
	if cfgFile != "" {
		options, err = utils.LoadConfiguration(cfgFile)
		if err != nil {
			log.Println("Error reading configuration file ", err)
			return err
		}
	}
	options.APIKey = os.Getenv("NASA_KEY")
	if options.APIKey == "" {
		log.Println("API for Nasa Services not set, please specify before continuing.")
		return fmt.Errorf("APIKey not Set")
	}

	if options.MaxPages == 0 {
		options.MaxPages = 10
	}

	if tleSatelite != -1 {
		var file []byte

		filename := fmt.Sprintf("%s.json", strconv.Itoa(tleSatelite))
		lteRet, err := explorer.GetTLEMember(options, tleSatelite)
		if err != nil {
			fmt.Println(err)
		}
		file, _ = json.MarshalIndent(lteRet, "", " ")
		_ = ioutil.WriteFile(filename, file, 0644)
	}

	for _, service := range services {
		var file []byte
		filename := fmt.Sprintf("%s.json", service)
		switch service {
		case "Apod":
			apodRet, err := explorer.GetLatestApod(options)
			if err != nil {
				fmt.Println(err)
			}
			file, _ = json.MarshalIndent(apodRet, "", " ")
			_ = ioutil.WriteFile(filename, file, 0644)

		case "EonetLatest":
			eonetRet, err := explorer.GetEonetLatestEvent(options)
			if err != nil {
				fmt.Println(err)
			}
			file, _ = json.MarshalIndent(eonetRet, "", " ")
			_ = ioutil.WriteFile(filename, file, 0644)

		case "NeoAll":
			neoRet, err := explorer.GetNeoAll(options)
			if err != nil {
				fmt.Println(err)
			}
			file, _ = json.MarshalIndent(neoRet, "", " ")
			_ = ioutil.WriteFile(filename, file, 0644)

		case "TLECollection":
			lteRet, err := explorer.GetAllTLECollection(options)
			if err != nil {
				fmt.Println(err)
			}
			file, _ = json.MarshalIndent(lteRet, "", " ")
			_ = ioutil.WriteFile(filename, file, 0644)
		}

	}

	return err
}
