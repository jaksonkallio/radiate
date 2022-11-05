package main

import (
	"log"

	"github.com/jaksonkallio/radiate/internal/media"
	"github.com/jaksonkallio/radiate/internal/service"

	ipfsapi "github.com/ipfs/go-ipfs-api"
)

func main() {
	service.InitLogger()

	err := media.InitDatabaseConnection()
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}

	// TODO: make this host string configurable.
	clientIPFS := ipfsapi.NewShell("localhost:5001")

	serviceInstance, err := service.NewService(clientIPFS)
	if err != nil {
		log.Fatalf("Could not start service: %s", err)
	}

	service.Serve()
}
