package media

import (
	"encoding/json"
)

// LibraryIndex is a serialized representation of a library and all the media it contains.
// Keep in mind that we can't fully "trust" the data in this file, because it's arbitrarily constructed by the library maintainer.
type LibraryIndex struct {
	SchemaVersion int                 `json:"schema_version"`
	Info          LibraryIndexInfo    `json:"info"`
	Announcement  LibraryAnnouncement `json:"announcement"`
	Items         []LibraryIndexMedia `json:"media"`
}

// LibraryIndexInfo is meta info about the library.
type LibraryIndexInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type LibraryAnnouncement struct {
	Text      string `json:"text"`
	UpdatedAt int    `json:"updated_at"`
}

type LibraryIndexMedia struct {
	Title         string `json:"title"`
	Size          int    `json:"size"`
	FileExtension string `json:"file_extension"`
	CID           CID    `json:"cid"`
	Description   string `json:"description"`
	AddedAt       int    `json:"added_at"`
}

func ParseLibraryIndexFile(indexFileBytes []byte) (LibraryIndex, error) {
	libraryIndex := LibraryIndex{}
	err := json.Unmarshal(indexFileBytes, &libraryIndex)
	if err != nil {
		return libraryIndex, err
	}

	return libraryIndex, nil
}
