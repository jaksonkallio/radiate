// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Library struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Motd        *string `json:"motd"`
	IpnsID      string  `json:"ipns_id"`
}

type Media struct {
	ID          string   `json:"id"`
	IpfsCid     string   `json:"ipfs_cid"`
	Extension   string   `json:"extension"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Starred     bool     `json:"starred"`
	Pinned      bool     `json:"pinned"`
}
