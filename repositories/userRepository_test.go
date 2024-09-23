package repository_test

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	// SQLite in-memory database for testing
	return gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
}
