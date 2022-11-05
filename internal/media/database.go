package media

import (
	"fmt"
	"github.com/rs/zerolog/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DatabaseConn *gorm.DB

var DatabaseModels = []interface{}{
	&Library{},
}

func InitDatabaseConnection() error {
	log.Info().Msg("initializing database connection")

	databaseConnection, err := gorm.Open(sqlite.Open("../data.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %s", err)
	}

	DatabaseConn = databaseConnection

	err = migrateSchemas()
	if err != nil {
		return err
	}

	return nil
}

func migrateSchemas() error {
	for _, databaseModel := range DatabaseModels {
		err := DatabaseConn.AutoMigrate(databaseModel)
		if err != nil {
			return err
		}
	}

	return nil
}
