package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"epyphite/space/v1/NASA/cmd/commands"
)

func main() {
	Environment := os.Getenv("EPY_SPACE_ENV")
	if Environment == "production" {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)

	} else {
		// The TextFormatter is default, you don't actually have to do this.
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.DebugLevel)
	}
	commands.Execute()
}
