package media

import (
	"sync"
	"time"

	"gorm.io/gorm"
)

type Library struct {
	gorm.Model
	Name              string    `gorm:"name"`
	Description       string    `gorm:"description"`
	MOTD              string    `gorm:"motd"`
	IPNSID            string    `gorm:"ipns_id"`
	InitialIngest     bool      `gorm:"initial_ingest"`
	LastIngestIPFSCID string    `gorm:"last_ingest_ipfs_cid"`
	LastIngested      time.Time `gorm:"last_ingested"`
	LastUpdated       time.Time `gorm:"last_updated"`
}

var LibraryIngesting map[uint]bool = make(map[uint]bool, 0)
var LibraryIngestingModMutex *sync.Mutex

// Re-ingests a library, updated the "last updated" time if anything has changed.
func (library *Library) Ingest() error {
	// Check whether this library is currently ingesting.
	LibraryIngestingModMutex.Lock()
	if library.CurrentlyIngesting() {
		return nil
	}

	// Mark ingesting to true.
	LibraryIngesting[library.ID] = true
	LibraryIngestingModMutex.Unlock()

	// Ingesting this library complete, mark as no longer currently ingesting.
	LibraryIngesting[library.ID] = false

	return nil
}

func (library *Library) FetchIndex() (*LibraryIndex, error) {
	return nil, nil
}

// Tells us whether a library is currently in the process of ingesting.
func (library *Library) CurrentlyIngesting() bool {
	ingesting, exists := LibraryIngesting[library.ID]
	return exists && ingesting
}

func FindLibraryByIPNSID(IPNSID string) (*Library, error) {
	library := []*Library{}

	tx := DatabaseConn.Limit(1).Where("ipns_id = ?", IPNSID).Find(&library)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if len(library) == 0 {
		return nil, nil
	}

	return library[0], nil
}

func FindOrInitializeLibraryByIPNSID(IPNSID string) (*Library, error) {
	library, err := FindLibraryByIPNSID(IPNSID)
	if err != nil {
		return nil, err
	}

	if library == nil {
		// Library not found, create.
		library = &Library{
			InitialIngest: false,
		}
	}

	return library, nil
}
