package media

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Extension         string    `gorm:"extension"`
	IPFSCID           string    `gorm:"ipfs_cid"`
	Description       string    `gorm:"description"`
	Starred           bool      `gorm:"starred"`
	Pinned            bool      `gorm:"pinned"`
	DateAdded         time.Time `gorm:"date_added"`
	DateFirstIngested time.Time `gorm:"date_first_ingested"`
	LibraryID         uint      `gorm:"library_id"`
	Library           Library   `gorm:"foreignKey:library_id"`
}
