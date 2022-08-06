package ipfs_client

import "fmt"

type ClientIPFS struct {
	Host string
}

func NewClientIPFS(host string) (*ClientIPFS, error) {
	clientIPFS := &ClientIPFS{
		Host: host,
	}

	if !clientIPFS.ServiceAvailable() {
		return nil, fmt.Errorf("could not reach IPFS service")
	}

	return clientIPFS, nil
}

func (clientIPFS *ClientIPFS) ServiceAvailable() bool {
	_, err := clientIPFS.Version()
	return err == nil
}
