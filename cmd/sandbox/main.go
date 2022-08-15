package main

import (
	"log"

	ipfsapi "github.com/ipfs/go-ipfs-api"
)

func main() {
	log.Println("Starting sandbox")
	/*
		ipfsAPI, err := ipfsapi.NewLocalApi()
		if err != nil {
			log.Fatalf("Failed to create API: %s", err)
		}

		ipfsNameAPI := ipfsAPI.Name()

		path, err := ipfsNameAPI.Resolve(context.Background(), "k51qzi5uqu5di69mv2h3e68jee8molk2i0u5m1h01nr7f1xsuhznl2lk7vrsum")
		if err != nil {
			log.Fatalf("Failed to resolve: %s", err)
		}

		log.Printf("Resolved path: %s", path)*/
	/*
		clientIPFS, err := ipfs_client.NewClientIPFS("localhost:5001")
		if err != nil {
			log.Fatalf("Could not create IPFS client: %s", err)
		}

		versionResponse, err := clientIPFS.Version()
		if err != nil {
			log.Fatalf("Could not get version: %s", err)
		}*/

	// Where your local node is running on localhost:5001
	clientIPFS := ipfsapi.NewShell("localhost:5001")

	err := clientIPFS.Get("QmbBnZBudfuGbbzDRTAhSVEYsoxCAe5yBimuuYvYu3BMHG", "/Users/jak/code/radiate/cache/")
	if err != nil {
		log.Fatalf("err: %s", err)
	}
}
