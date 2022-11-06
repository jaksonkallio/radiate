package media

import (
	"encoding/json"
	"fmt"
	"github.com/c2h5oh/datasize"
	"github.com/google/uuid"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/jaksonkallio/radiate/internal/config"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"path"
	"time"

	"gorm.io/gorm"
)

const IngestMaxIndexSize = 10 * datasize.MB

type Library struct {
	gorm.Model
	UniqueIdentifier      string          `gorm:"column:unique_identifier;unique"`
	IngestedTitle         string          `gorm:"ingested_title"`
	Description           string          `gorm:"description"`
	Announcement          string          `gorm:"column:announcement"`
	AnnouncementUpdatedAt time.Time       `gorm:"column:announcement_updated_at"`
	IgnoreAnnouncements   bool            `gorm:"column:ignore_announcements"`
	IndexIdentifier       IndexIdentifier `gorm:"column:index_identifier;unique"`
	InitialIngest         bool            `gorm:"initial_ingest"`
	IngestedAt            time.Time       `gorm:"ingested_at"`
	NewMediaAt            time.Time       `gorm:"new_media_at"`
}

// Ingest will re-ingest a library, updated the "last updated" time if anything has changed.
func (library *Library) Ingest(shellIPFS *shell.Shell) error {
	indexCID, err := library.IndexIdentifier.ResolveToCID()
	if err != nil {
		return errors.Wrap(err, "could not resolve CID")
	}

	indexFileLocation := path.Join(
		config.CurrentConfig.CacheDir,
		"library_index",
		fmt.Sprintf("%s.json", library.UniqueIdentifier),
	)

	objectStats, err := shellIPFS.ObjectStat(indexCID.String())
	if err != nil {
		return err
	}

	indexFileCumulativeSize := datasize.ByteSize(objectStats.CumulativeSize)

	// Protect against parsing a massive JSON index file.
	if indexFileCumulativeSize > IngestMaxIndexSize {
		return fmt.Errorf("index file is too big, must be no more than %s", IngestMaxIndexSize.HumanReadable())
	}

	// Mark ingest time and download the index file locally.
	now := time.Now()
	err = shellIPFS.Get(indexCID.String(), indexFileLocation)
	if err != nil {
		return errors.Wrap(err, "get library index file failed")
	}

	log.Debug().
		Str("indexFileLocation", indexFileLocation).
		Str("size", indexFileCumulativeSize.HumanReadable()).
		Msg("downloaded index file")

	// TODO (@JaksonKallio): support different index file schema versions for backwards compatibility.

	// Now read the downloaded index file.
	indexFileBytes, err := ioutil.ReadFile(indexFileLocation)
	libraryIndex := Index{}
	err = json.Unmarshal(indexFileBytes, &libraryIndex)
	if err != nil {
		return errors.Wrap(err, "could not parse library index file")
	}

	// Update basic library info.
	library.Announcement = libraryIndex.Announcement.Text
	library.AnnouncementUpdatedAt = time.Unix(libraryIndex.Announcement.UpdatedAt, 0)
	library.Description = libraryIndex.Info.Description
	library.IngestedTitle = libraryIndex.Info.Title

	// Update fields related to ingest status.
	library.IngestedAt = now
	library.InitialIngest = true

	// Add/update any media.
	var newItemCount int
	var ingestedItemCount int
	for _, media := range libraryIndex.Medias {
		if !media.CID.Valid() {
			continue
		}

		item, err := FindItemByCID(media.CID)
		if err != nil {
			return err
		}

		if item == nil {
			item, err = CreateItem(library.ID, media.CID)
			if err != nil {
				return err
			}

			// Mark the item's first-ingested time.
			item.FirstIngestedAt = now
			newItemCount += 1
		}

		ingestedItemCount += 1

		// Update item info.
		item.Description = media.Description
		item.AddedAt = time.Unix(media.AddedAt, 0)
		item.FileExtension = media.FileExtension

		// Persist media item.
		DatabaseConn.Save(item)

		go func() {
			err := item.RefreshObjectStats(shellIPFS)
			if err != nil {
				log.Warn().
					Err(err).
					Msg("could not refresh media item object stats after ingest")
			}
		}()
	}

	log.Debug().
		Int("ingestedItemCount", ingestedItemCount).
		Int("newItemCount", newItemCount).
		Msg("ingested media items")

	if newItemCount > 0 {
		library.NewMediaAt = now
	}

	// Save updated library info.
	DatabaseConn.Save(library)

	return nil
}

func FindLibraryByID(libraryID uint) (library *Library, err error) {
	library = &Library{}

	err = DatabaseConn.Find(library, libraryID).Error
	if err != nil {
		return nil, err
	}

	return
}

func CreateLibraryByIndexIdentifier(indexIdentifier IndexIdentifier) (*Library, error) {
	countWithIndexIdentifier, err := CountLibrariesWithIndexIdentifier(indexIdentifier)
	if err != nil {
		return nil, errors.Wrap(err, "could not count number of libraries with index identifier")
	}

	if countWithIndexIdentifier > 0 {
		return nil, fmt.Errorf("a library with index identifier %q already exists", indexIdentifier.ValueString)
	}

	libraryUUID := uuid.New()
	library := &Library{
		UniqueIdentifier: libraryUUID.String(),
		IndexIdentifier:  indexIdentifier,
	}

	err = DatabaseConn.Create(library).Error
	if err != nil {
		return nil, errors.Wrap(err, "could not create library")
	}

	return library, nil
}

func CountLibrariesWithIndexIdentifier(indexIdentifier IndexIdentifier) (count int64, err error) {
	err = DatabaseConn.Model(&Library{}).Where("index_identifier = ?", indexIdentifier).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return
}
