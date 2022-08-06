package main

import (
	"log"

	"github.com/jaksonkallio/radiate/internal/media"
	"github.com/jaksonkallio/radiate/internal/service"
)

func main() {
	err := media.InitDatabaseConnection()
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}

	service, err := service.NewService()
	if err != nil {
		log.Fatalf("Could not start service: %s", err)
	}

	err = media.InitIPFSClient()
	if err != nil {
		log.Fatalf("Could not initialize IPFS client: %s", err)
	}

	service.Serve()
}
