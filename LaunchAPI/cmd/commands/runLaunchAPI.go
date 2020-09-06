package commands

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/epyphite/space/LaunchAPI/pkg/constants"
	"github.com/epyphite/space/LaunchAPI/pkg/models"
	"github.com/epyphite/space/LaunchAPI/pkg/utils"
	web "github.com/epyphite/space/LaunchAPI/pkg/web"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "launchAPI",
	Short: "launchAPI",
	Long:  "",
	RunE:  launchAPI,
}
var cfgFile string
var services []string
var tleSatelite int
var clearDB bool
var webServer bool

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Define a configuration file location")
	rootCmd.PersistentFlags().StringArrayVar(&services, "services", nil, "Define the services you wish to run")
	rootCmd.PersistentFlags().BoolVar(&clearDB, "cleardb", false, "Clear db will clear the DB upon start")
}

//Execute will run the desire module command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func launchAPI(cmd *cobra.Command, args []string) error {

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

	options.WebAddress = constants.Webaddress
	options.WebPort = constants.Webport
	options.DatabaseName = constants.Databasename
	if clearDB == true {
		err := os.Remove(constants.ContentDir + options.DatabaseName)
		if err != nil {
			log.Errorln("Could not delete database ", err)
		}
	}

	webagent, err := web.NewWebAgent(options)
	if err != nil {
		log.Fatalln("Error on newebagent call ", err)
	}
	webagent.StartServer()
	return err
}
