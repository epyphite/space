package esaExplorer

import (
	"context"
	"os"

	client "epyphite/space/v1/ESA/pkg/client"
	models "epyphite/space/v1/ESA/pkg/models"
	modules "epyphite/space/v1/ESA/pkg/models/modules"
)

func GetDISCUSALL(configuration models.Config) ([]*modules.DiscosResponse, error) {
	var config models.Config
	config = configuration
	if config == (models.Config{}) {
		config = models.Config{APIKey: os.Getenv("ESA_KEY")}
	}

	if config.BaseURL == "" {
		config.BaseURL = "https://discosweb.esoc.esa.int"
	}
	DiscusOptions := modules.DiscosRequest{}

	c := client.NewClient(config)
	ctx := context.Background()
	res, err := c.GetDiscusObjects(ctx, &DiscusOptions)
	return res, err
}
