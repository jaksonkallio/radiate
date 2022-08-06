package ipfs_client

import "encoding/json"

type VersionResponse struct {
	Commit  string
	Golang  string
	Repo    string
	System  string
	Version string
}

func (clientIPFS *ClientIPFS) Version() (*VersionResponse, error) {
	data, err := clientIPFS.request("version", map[string]string{})
	if err != nil {
		return nil, err
	}

	versionResponse := &VersionResponse{}
	json.Unmarshal(data, versionResponse)

	return versionResponse, nil
}
