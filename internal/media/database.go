package media

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DatabaseConn *gorm.DB

func InitDatabaseConnection() error {
	databaseConnection, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %s", err)
	}

	DatabaseConn = databaseConnection

	migrateSchemas()

	return nil
}

func migrateSchemas() {
	DatabaseConn.AutoMigrate(&Library{})
}
