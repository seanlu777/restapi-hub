// db/db.go
package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Initialize connects to the PostgreSQL and runs migrations
func Initialize(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to db: %v", err)
	}

	err = DB.AutoMigrate(&HubHistory{}, A2TB{}, R2B2{}, R2T8{})
	if err != nil {
		return fmt.Errorf("failed to migrate db: %v", err)
	}

	return nil
}
