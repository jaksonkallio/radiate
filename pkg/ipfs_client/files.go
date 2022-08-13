package ipfs_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type Entry struct {
	Name string `json:"Name"`
	Type int    `json:"Type"`
	Size int    `json:"Size"`
	Hash string `json:"Hash"`
}

type Ls struct {
	Entries []Entry `json:"Entries"`
}

func (clientIPFS *ClientIPFS) Ls(path string, long bool, directoryOrder bool) (*Ls, error) {
	data, err := clientIPFS.request(
		"ls",
		map[string]string{
			"arg":  path,
			"long": BoolString(long),
			"U":    BoolString(directoryOrder),
		},
		&bytes.Buffer{},
	)
	if err != nil {
		return nil, err
	}

	lsResponse := &Ls{}
	json.Unmarshal(data, lsResponse)

	return lsResponse, nil
}

func (clientIPFS *ClientIPFS) MkDir(path string, createParents bool, cidVersion int, hash string) error {
	params := map[string]string{
		"arg":     path,
		"parents": BoolString(createParents),
	}

	if cidVersion >= 0 {
		params["cid-version"] = fmt.Sprintf("%d", cidVersion)
	}

	if len(hash) > 0 {
		params["hash"] = hash
	}

	_, err := clientIPFS.request("mkdir", params, &bytes.Buffer{})
	if err != nil {
		return err
	}

	return nil
}

func (clientIPFS *ClientIPFS) Write(
	localFilePath string,
	destinationPath string,
	beginWriteOffset int64,
	createIfNotExists bool,
	createParents bool,
	truncateBeforeWrite bool,
	maxBytesRead int64,
	rawBlocksNewLeafNodes bool,
	cidVersion int,
	hash string,
) error {
	params := map[string]string{
		"arg":        destinationPath,
		"create":     BoolString(createIfNotExists),
		"parents":    BoolString(createParents),
		"truncate":   BoolString(truncateBeforeWrite),
		"raw-leaves": BoolString(rawBlocksNewLeafNodes),
	}

	if beginWriteOffset >= 0 {
		params["offset"] = fmt.Sprintf("%d", beginWriteOffset)
	}

	if maxBytesRead >= 0 {
		params["count"] = fmt.Sprintf("%d", maxBytesRead)
	}

	if cidVersion >= 0 {
		params["cid-version"] = fmt.Sprintf("%d", cidVersion)
	}

	if len(hash) > 0 {
		params["hash"] = hash
	}

	// Open the local file.
	file, err := os.Open(localFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", localFilePath)
	if err != nil {
		return err
	}

	io.Copy(part, file)
	writer.Close()

	_, err = clientIPFS.request("write", params, body)
	if err != nil {
		return err
	}

	return nil
}

func (clientIPFS *ClientIPFS) Read(
	path string,
	beginReadOffset int64,
	maxBytesRead int64,
) ([]byte, error) {
	params := map[string]string{
		"arg": path,
	}

	if beginReadOffset >= 0 {
		params["offset"] = fmt.Sprintf("%d", beginReadOffset)
	}

	if maxBytesRead >= 0 {
		params["count"] = fmt.Sprintf("%d", maxBytesRead)
	}

	data, err := clientIPFS.request("read", params, &bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	return data, nil
}
