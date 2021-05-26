package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"epyphite/space/SatCom/pkg/models"
	utils "epyphite/space/SatCom/pkg/utils"
)

var rootCmd = &cobra.Command{
	Use:   "satCom",
	Short: "satCom",
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config/SatCom.json)")

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

	log.Println("Starting SatCom Utilities ", config.Version)

	go utils.HandleSignal()
	return err
}
