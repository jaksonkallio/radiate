package media

// Index is a serialized representation of a library and all the media it contains.
// Keep in mind that we can't fully "trust" the data in this file, because it's arbitrarily constructed by the library maintainer.
type Index struct {
	SchemaVersion int               `json:"schema_version"`
	Info          IndexInfo         `json:"info"`
	Announcement  IndexAnnouncement `json:"announcement"`
	Medias        []IndexMedia      `json:"medias"`
}

// IndexInfo is meta info about the library.
type IndexInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type IndexAnnouncement struct {
	Text      string `json:"text"`
	UpdatedAt int64  `json:"updated_at"`
}

type IndexMedia struct {
	Title         string `json:"title"`
	FileExtension string `json:"file_extension"`
	CID           CID    `json:"cid"`
	Description   string `json:"description"`
	AddedAt       int64  `json:"added_at"`
}
