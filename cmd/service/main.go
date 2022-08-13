package main

import (
	"log"

	"github.com/jaksonkallio/radiate/internal/media"
	"github.com/jaksonkallio/radiate/internal/service"
	"github.com/jaksonkallio/radiate/pkg/ipfs_client"
)

func main() {
	err := media.InitDatabaseConnection()
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}

	// TODO: make this host string configurable.
	clientIPFS, err := ipfs_client.NewClientIPFS("localhost:5001")
	if err != nil {
		log.Fatalf("Could not create IPFS client: %s", err)
	}

	service, err := service.NewService(clientIPFS)
	if err != nil {
		log.Fatalf("Could not start service: %s", err)
	}

	service.Serve()
}
