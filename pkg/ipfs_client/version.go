package ipfs_client

import (
	"bytes"
	"encoding/json"
)

type Version struct {
	Commit  string
	Golang  string
	Repo    string
	System  string
	Version string
}

func (clientIPFS *ClientIPFS) Version() (*Version, error) {
	data, err := clientIPFS.request("version", map[string]string{}, &bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	versionResponse := &Version{}
	json.Unmarshal(data, versionResponse)

	return versionResponse, nil
}
