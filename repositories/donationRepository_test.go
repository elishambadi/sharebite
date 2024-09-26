package repository_test

import (
	"testing"

	"github.com/elishambadi/sharebite/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndFindAll(t *testing.T) {
	db, _ := setupTestDB()
	db.AutoMigrate(&models.User{}, &models.Donation{}, &models.DonationRequest{})

	fakeUsers := models.FakeUsers()
	for _, user := range fakeUsers {
		db.Create(&user)
	}

	donations := models.GenerateFakeDonations()

	result1 := db.Create(donations[0])
	result2 := db.Create(donations[1])

	assert.NoError(t, result1.Error)
	assert.NoError(t, result2.Error)

	// Find all items
	var donationsFromDB []models.Donation
	err := db.Find(&donationsFromDB).Error

	assert.NoError(t, err)
	assert.Equal(t, 2, len(donationsFromDB))
}
