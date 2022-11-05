package main

import (
	"github.com/jaksonkallio/radiate/internal/config"
	"github.com/rs/zerolog/log"

	"github.com/jaksonkallio/radiate/internal/media"
	"github.com/jaksonkallio/radiate/internal/service"
)

func main() {
	// Load the config.
	err := config.LoadFromFile()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("could not load config")
	}

	service.InitLogger()

	// Initialize our database connection.
	err = media.InitDatabaseConnection()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("could not initialize database connection")
	}

	// Create the service instance.
	serviceInstance, err := service.NewService()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("could not start service")
	}

	serviceInstance.ServeAPI()
}
