package repository_test

import (
	"testing"

	"github.com/elishambadi/sharebite/models"
	repository "github.com/elishambadi/sharebite/repositories"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateAndFindDonationAndRequests(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	db.AutoMigrate(&models.User{}, &models.Donation{}, &models.DonationRequest{})

	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)
	userRepo := repository.NewGormUserRepository(db, logger)

	fakeUsers := models.FakeUsers()
	for _, user := range fakeUsers {
		userRepo.Create(user)
	}

	donationRepo := repository.NewDonationRepository(db)
	donations := models.GenerateFakeDonations()
	fakeRequests := models.GenerateFakeDonationRequests()

	for _, donation := range donations {
		err := donationRepo.CreateDonation(&donation)
		assert.NoError(t, err)
	}

	// Find all items
	donationsFromDB, err := donationRepo.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(donationsFromDB))

	t.Run("test donation requests", func(t *testing.T) {
		for _, request := range fakeRequests {
			err := donationRepo.CreateDonationRequest(&request)
			assert.NoError(t, err)
		}

		// List requests
		requests, err := donationRepo.ListDonationRequests()
		assert.NoError(t, err)
		assert.Equal(t, 3, len(requests))

		// Get request by ID
		var requestFromDb *models.DonationRequest
		requestFromDb, err = donationRepo.GetDonationRequestByID("1")
		assert.NoError(t, err)
		assert.Equal(t, "Pending", requestFromDb.Status)

		// Update request status
		updateErr := donationRepo.UpdateDonationRequestStatus("1", "Approved")
		assert.NoError(t, updateErr)
		requestFromDb, err = donationRepo.GetDonationRequestByID("1")
		assert.NoError(t, err)
		assert.Equal(t, "Approved", requestFromDb.Status)
	})
}
