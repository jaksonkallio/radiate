package ipfs_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ErrorIPFS struct {
	Message string
	Code    int
	Type    string
}

func (errorIPFS *ErrorIPFS) Error() string {
	return errorIPFS.Message
}

func (clientIPFS *ClientIPFS) request(endpoint string, query map[string]string, body *bytes.Buffer) ([]byte, error) {
	urlValues := url.Values{}

	for key, val := range query {
		urlValues.Set(key, val)
	}

	res, err := http.Post(fmt.Sprintf("http://%s/api/v0/%s?%s", clientIPFS.Host, endpoint, urlValues.Encode()), "application/json", body)
	if err != nil {
		return nil, err
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		err = &ErrorIPFS{}

		parseErrorErr := json.Unmarshal(resData, err)
		if parseErrorErr != nil {
			return nil, fmt.Errorf("Received status code %d from IPFS server, but could not parse the error message received", res.StatusCode)
		}

		return nil, err
	}

	return resData, nil
}
