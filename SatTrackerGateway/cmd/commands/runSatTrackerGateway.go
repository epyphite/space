package commands

import (
	"encoding/json"
	"epyphite/space/v1/SatTrackerGateway/pkg/models"
	"epyphite/space/v1/SatTrackerGateway/srv"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "satExplorer",
	Short: "satExplorer",
	Long:  ``,
	RunE:  satExplorer,
}
var cfgFile string
var webServer bool

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Define a configuration file location")
	rootCmd.PersistentFlags().BoolVar(&webServer, "event", false, "Event Tracker")

}

//Execute will run the desire module command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func satExplorer(cmd *cobra.Command, args []string) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	satExplorer := srv.SatTrackAPI{}
	identity := models.Identity{}
	identity.Identity = username
	identity.Password = password

	_, err = satExplorer.Login(identity)
	if err != nil {
		log.Println("Error ", err)
	}
	ret1, err := satExplorer.MakeRequestBoxScore()
	b, err := json.MarshalIndent(ret1, "", "\t")
	if err != nil {
		log.Println("Error ", err)
	}
	log.Println(string(b))
	return err

}
