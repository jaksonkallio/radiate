package media

import (
	ipfsapi "github.com/ipfs/go-ipfs-api"
	"sync"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	FileExtension   string     `gorm:"column:file_extension"`
	Size            int        `gorm:"column:size"`
	CID             string     `gorm:"column:cid;unique"`
	Description     string     `gorm:"column:description"`
	Starred         bool       `gorm:"column:starred"`
	Pinned          bool       `gorm:"column:pinned"`
	AddedAt         time.Time  `gorm:"column:added_at"`
	FirstIngestedAt time.Time  `gorm:"column:first_ingested_at"`
	LibraryID       uint       `gorm:"column:library_id"`
	Library         Library    `gorm:"foreignKey:library_id"`
	objectStatsLock sync.Mutex `gorm:"-"`
}

func FindItemByCID(cID CID) (item *Item, err error) {
	item = &Item{}

	result := DatabaseConn.Where("cid = ?", cID.String()).Limit(1).Find(item)
	err = result.Error
	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		item = nil
	}

	return
}

func CreateItem(libraryID uint, cID CID) (item *Item, err error) {
	item = &Item{
		CID:       cID.String(),
		LibraryID: libraryID,
	}

	err = DatabaseConn.Create(item).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (item *Item) RefreshObjectStats(shellIPFS *ipfsapi.Shell) error {
	objectStats, err := shellIPFS.ObjectStat(item.CID)
	if err != nil {
		return err
	}

	item.Size = objectStats.CumulativeSize

	// TODO (@JaksonKallio): consider saving other IPFS object stats here.

	return DatabaseConn.Save(item).Error
}
