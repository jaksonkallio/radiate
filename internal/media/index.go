package media

import (
	"encoding/json"
	"time"
)

// An "index" is a serialized representation of a library and all of the media it contains.
type LibraryIndex struct {
	SchemaVersion int                `json:"schema_version"`
	Info          LibraryIndexInfo   `json:"info"`
	Items         []LibraryIndexItem `json:"items"`
}

// Meta info about the library.
type LibraryIndexInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	MOTD        string `json:"motd"`
}

type LibraryIndexItem struct {
	Extension   string    `json:"extension"`
	IPFSCID     string    `json:"ipfs_cid"`
	Description string    `json:"description"`
	DateAdded   time.Time `json:"date_added"`
}

func ParseLibraryIndexFile(indexFileBytes []byte) (LibraryIndex, error) {
	libraryIndex := LibraryIndex{}
	err := json.Unmarshal(indexFileBytes, &libraryIndex)
	if err != nil {
		return libraryIndex, err
	}

	return libraryIndex, nil
}
