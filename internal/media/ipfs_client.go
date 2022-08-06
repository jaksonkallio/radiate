package media

import (
	"log"

	ipfsapi "github.com/ipfs/go-ipfs-http-client"
)

var IPFSClient *ipfsapi.HttpApi

func InitIPFSClient() error {
	clientIPFS, err := ipfsapi.NewLocalApi()
	if err != nil {
		log.Fatalf("Failed to create IPFS API client: %s", err)
	}

	IPFSClient = clientIPFS

	return nil
}
