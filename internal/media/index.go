package media

import (
	"encoding/json"
	"time"
)

// LibraryIndex is a serialized representation of a library and all the media it contains.
// Keep in mind that we can't fully "trust" the data in this file, because it's arbitrarily constructed by the library maintainer.
type LibraryIndex struct {
	SchemaVersion int                 `json:"schema_version"`
	Info          LibraryIndexInfo    `json:"info"`
	Items         []LibraryIndexMedia `json:"media"`
}

// LibraryIndexInfo is meta info about the library.
type LibraryIndexInfo struct {
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	IPFSCID          string    `json:"ipfs_cid"`
	MOTD             string    `json:"motd"`
	LastChangedMOTD  time.Time `json:"last_changed_motd"`
	LastChangedMedia time.Time `json:"last_changed_media"`
}

type LibraryIndexMedia struct {
	Title         string    `json:"title"`
	FileExtension string    `json:"file_extension"`
	IPFSCID       string    `json:"ipfs_cid"`
	Description   string    `json:"description"`
	DateAdded     time.Time `json:"date_added"`
}

func ParseLibraryIndexFile(indexFileBytes []byte) (LibraryIndex, error) {
	libraryIndex := LibraryIndex{}
	err := json.Unmarshal(indexFileBytes, &libraryIndex)
	if err != nil {
		return libraryIndex, err
	}

	return libraryIndex, nil
}
