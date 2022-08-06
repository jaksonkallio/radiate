package ipfs_client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (clientIPFS *ClientIPFS) request(endpoint string, query map[string]string) ([]byte, error) {
	urlValues := url.Values{}

	for key, val := range query {
		urlValues.Set(key, val)
	}

	res, err := http.Post(fmt.Sprintf("http://%s/api/v0/%s?%s", clientIPFS.Host, endpoint, urlValues.Encode()), "application/json", bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resData, nil
}
