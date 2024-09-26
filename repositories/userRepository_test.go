package repository_test

import (
	"testing"

	"github.com/elishambadi/sharebite/models"
	repository "github.com/elishambadi/sharebite/repositories"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	// SQLite in-memory database for testing
	return gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
}

func TestCreateAndFindAll(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	db.AutoMigrate(&models.User{})

	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)
	userRepo := repository.NewGormUserRepository(db, logger)

	fakeUsers := models.FakeUsers()

	for _, user := range fakeUsers {
		userRepo.Create(user)
	}

	users, err := userRepo.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, fakeUsers[0].Email, users[0].Email)
}

func TestGetUserByEmail(t *testing.T) {
	// Setup
	db, err := setupTestDB()
	assert.NoError(t, err)

	db.AutoMigrate(&models.User{})

	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)
	userRepo := repository.NewGormUserRepository(db, logger)

	fakeUsers := models.FakeUsers()

	for _, user := range fakeUsers {
		userRepo.Create(user)
	}

	// Tests
	user, err := userRepo.GetUserByEmail("bob@example.com")
	assert.Equal(t, "bob@example.com", user.Email)
	assert.NoError(t, err)

	t.Run("user not found", func(t *testing.T) {
		user, err := userRepo.GetUserByEmail("invalid@example.com")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
		assert.Equal(t, models.User{}, user) // Assert user is an empty struct
	})
}

func TestGetUserById(t *testing.T) {
	// Setup
	db, err := setupTestDB()
	assert.NoError(t, err)

	db.AutoMigrate(&models.User{})

	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)
	userRepo := repository.NewGormUserRepository(db, logger)

	fakeUsers := models.FakeUsers()

	for _, user := range fakeUsers {
		userRepo.Create(user)
	}

	// Test itself
	user, err := userRepo.GetUserById("2")
	assert.Equal(t, "bob@example.com", user.Email)
	assert.NoError(t, err)

	t.Run("user not found", func(t *testing.T) {
		user, err := userRepo.GetUserById("9999999999")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
		assert.Equal(t, models.User{}, user) // Assert user is an empty struct
	})
}

func TestUpdateAPIToken(t *testing.T) {
	// Setup
	db, err := setupTestDB()
	assert.NoError(t, err)

	db.AutoMigrate(&models.User{})

	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)
	userRepo := repository.NewGormUserRepository(db, logger)

	fakeUsers := models.FakeUsers()

	for _, user := range fakeUsers {
		userRepo.Create(user)
	}

	// Test itself
	user, _ := userRepo.GetUserById("1") // Get the first user
	updateErr := userRepo.UpdateAPIToken(&user, "sample-api-token")
	assert.NoError(t, updateErr)
	assert.Equal(t, "sample-api-token", user.APIToken)
}

func TestDeleteUserById(t *testing.T) {
	//Setup
	db, err := setupTestDB()
	assert.NoError(t, err)

	db.AutoMigrate(&models.User{})

	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)
	userRepo := repository.NewGormUserRepository(db, logger)

	fakeUsers := models.FakeUsers()

	for _, user := range fakeUsers {
		userRepo.Create(user)
	}

	// Test
	deleteErr := userRepo.DeleteUserById("1")
	assert.NoError(t, deleteErr)

	allUsers, err := userRepo.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(allUsers))
}
