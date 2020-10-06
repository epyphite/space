package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"epyphite/space/v1/LaunchWeb/pkg/models"
	"epyphite/space/v1/LaunchWeb/pkg/web"
	"github.com/spf13/cobra"

	utils "epyphite/space/v1/LaunchWeb/pkg/utils"
)

var rootCmd = &cobra.Command{
	Use:   "resilientApi",
	Short: "resilientApi",
	Long:  ``,
	RunE:  runWeb,
}

//Execute will run the desire module command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var config models.Config
var cfgFile string
var projectBase string
var monitorMode string

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config/apiLayer.json)")

}

func runWeb(cmd *cobra.Command, args []string) error {
	execName, err := os.Executable()
	baseDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	execName = filepath.Base(execName)

	if cfgFile == "" {
		config, _ = utils.LoadConfigurationDefaults()
		cfgFile = baseDir + config.TempDir + execName + ".json"
	}

	config, err := utils.LoadConfiguration(cfgFile)
	if err != nil {
		log.Fatalln("Error and Exiting ")
	}

	webagent, err := web.NewWebAgent(config)
	if err != nil {
		log.Fatalln("Error on newebagent call ", err)
	}
	log.Println("Starting Web Server in", config.WebAddress, config.WebPort)

	go utils.HandleSignal()
	webagent.StartServer()
	return err
}
