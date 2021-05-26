package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	explorer "epyphite/space/v1/ESA"
	"epyphite/space/v1/ESA/pkg/constants"
	"epyphite/space/v1/ESA/pkg/models"
	"epyphite/space/v1/ESA/pkg/utils"
	webapi "epyphite/space/v1/ESA/pkg/web"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "esaExplorer",
	Short: "esaExplorer",
	Long:  ``,
	RunE:  esaExplorer,
}
var cfgFile string
var services []string
var webServer bool

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Define a configuration file location")
	rootCmd.PersistentFlags().StringArrayVar(&services, "services", nil, "Define the services you wish to run")
	rootCmd.PersistentFlags().BoolVar(&webServer, "webServer", false, "Start Web Server")

}

//Execute will run the desire module command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func esaExplorer(cmd *cobra.Command, args []string) error {

	execName, _ := os.Executable()
	execName = filepath.Base(execName)
	options := models.Config{}
	var err error
	if cfgFile != "" {
		options, err = utils.LoadConfiguration(cfgFile)
		if err != nil {
			log.Errorln("Error reading configuration file ", err)
			return err
		}
	}
	options.APIKey = os.Getenv("NASA_KEY")
	if options.APIKey == "" {
		log.Errorln("API for Nasa Services not set, please specify before continuing.")
		return fmt.Errorf("APIKey not Set")
	}

	if options.MaxPages == 0 {
		options.MaxPages = 10
	}
	options.WebAddress = constants.Webaddress
	options.WebPort = constants.Webport
	options.DatabaseName = constants.Databasename

	if webServer == true {
		webagent, err := webapi.NewWebAgent(options)
		if err != nil {
			log.Fatalln("Error on newebagent call ", err)
		}
		webagent.StartServer()
	} else {

		for _, service := range services {
			var file []byte
			filename := fmt.Sprintf("%s.json", service)
			switch service {
			case "DISCUS":
				apodRet, err := explorer.GetDISCUSALL(options)
				if err != nil {
					log.Errorln(err)
				}
				file, _ = json.MarshalIndent(apodRet, "", " ")
				_ = ioutil.WriteFile(filename, file, 0644)
			}

		}
	}

	return err
}
