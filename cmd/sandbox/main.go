package main

import (
	"context"
	"log"

	ipfsapi "github.com/ipfs/go-ipfs-http-client"
)

func main() {
	log.Println("Starting sandbox")

	ipfsAPI, err := ipfsapi.NewLocalApi()
	if err != nil {
		log.Fatalf("Failed to create API: %s", err)
	}

	ipfsNameAPI := ipfsAPI.Name()

	path, err := ipfsNameAPI.Resolve(context.Background(), "k51qzi5uqu5di69mv2h3e68jee8molk2i0u5m1h01nr7f1xsuhznl2lk7vrsum")
	if err != nil {
		log.Fatalf("Failed to resolve: %s", err)
	}

	log.Printf("Resolved path: %s", path)
}
