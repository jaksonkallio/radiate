package main

import (
	"github.com/jaksonkallio/radiate/internal/media"
	"github.com/jaksonkallio/radiate/internal/service"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

const (
	IndexIdentifierKey = "index_identifier"
	IDKey              = "id"
	IPFSCIDKey         = "ipfs_cid"
)

var Commands = []Command{
	{
		Name: "library/ingest",
		Arguments: []string{
			IDKey,
		},
		Fn: func(service *service.Service, args map[string]string) error {
			libraryID, err := strconv.ParseUint(args[IDKey], 10, 32)
			if err != nil {
				return errors.Wrap(err, "did not provide a valid library ID")
			}

			library, err := media.FindLibraryByID(uint(libraryID))
			if err != nil {
				return errors.Wrap(err, "could not find library by ID")
			}

			err = library.Ingest(service.ShellIPFS)
			if err != nil {
				return errors.Wrap(err, "could not ingest library")
			}

			return nil
		},
	},
	{
		Name: "library/add",
		Arguments: []string{
			IndexIdentifierKey,
		},
		Fn: func(service *service.Service, args map[string]string) error {
			indexIdentifier := args[IndexIdentifierKey]
			library, err := media.CreateLibraryByIndexIdentifier(media.NewIndexIdentifierFromString(indexIdentifier))
			if err != nil {
				return err
			}

			log.Info().
				Str("unique_identifier", library.UniqueIdentifier).
				Int("id", int(library.ID)).
				Msg("added library")

			return nil
		},
	},
	{
		Name: "ipfs/get_object_stats",
		Arguments: []string{
			IPFSCIDKey,
		},
		Fn: func(service *service.Service, args map[string]string) error {
			objectStats, err := service.ShellIPFS.ObjectStat(args[IPFSCIDKey])
			if err != nil {
				return err
			}

			log.Info().
				Str("hash", objectStats.Hash).
				Int("cumulative_size", objectStats.CumulativeSize).
				Int("block_size", objectStats.BlockSize).
				Int("data_size", objectStats.DataSize).
				Int("links_size", objectStats.LinksSize).
				Int("num_links", objectStats.NumLinks).
				Msg("get object stats")

			return nil
		},
	},
}
