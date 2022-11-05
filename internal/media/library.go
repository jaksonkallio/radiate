package media

import (
	shell "github.com/ipfs/go-ipfs-api"
	"time"

	"gorm.io/gorm"
)

type Library struct {
	gorm.Model
	Title           string          `gorm:"title"`
	Description     string          `gorm:"description"`
	MOTD            string          `gorm:"motd"`
	IndexIdentifier IndexIdentifier `gorm:"index_identifier"`
	InitialIngest   bool            `gorm:"initial_ingest"`
	IngestedAt      int             `gorm:"ingested_at"`
	MediaUpdatedAt  time.Time       `gorm:"last_updated"`
}

// Ingest will re-ingest a library, updated the "last updated" time if anything has changed.
func (library *Library) Ingest(shellIPFS *shell.Shell) {
	//indexCID, _ := library.IndexIdentifier.ResolveToCID()

	//shellIPFS.Get(indexCID, "")
}

func (library *Library) FetchIndex() (*LibraryIndex, error) {
	return nil, nil
}

// Finds a library using an index identifier.
func FindLibraryByIndexIPFSIdentifier(identifier string) (*Library, error) {
	library := []*Library{}

	tx := DatabaseConn.Limit(1).Where("index_identifier = ?", identifier).Find(&library)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if len(library) == 0 {
		return nil, nil
	}

	return library[0], nil
}

func FindOrInitializeLibraryByIPFSIdentifier(identifier string) (*Library, error) {
	library, err := FindLibraryByIndexIPFSIdentifier(identifier)
	if err != nil {
		return nil, err
	}

	if library == nil {
		// Library not found, create.
		library = &Library{}

	}

	return library, nil
}
