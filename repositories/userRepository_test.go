package repository_test

import (
	"testing"

	"github.com/elishambadi/sharebite/models"
	repository "github.com/elishambadi/sharebite/repositories"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	// SQLite in-memory database for testing
	return gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
}

func TestFindAllUsers(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	db.AutoMigrate(&models.User{})

	fakeUsers := models.FakeUsers()

	for _, user := range fakeUsers {
		db.Create(&user)
	}

	userRepo := repository.NewGormUserRepository(db)

	users, err := userRepo.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, fakeUsers[0].Email, users[0].Email)
}
