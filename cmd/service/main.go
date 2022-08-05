package main

import (
	"log"

	"github.com/jaksonkallio/radiate/internal/service"
)

func main() {
	service, err := service.NewService()
	if err != nil {
		log.Fatalf("Could not start service: %s", err)
	}

	service.Serve()
}
