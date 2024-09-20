package services

import (
	"strconv"

	"github.com/elishambadi/sharebite/models" // Adjust the import path as needed
	"github.com/elishambadi/sharebite/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DonationService struct {
	db *gorm.DB
}

func NewDonationService(db *gorm.DB) *DonationService {
	return &DonationService{db: db}
}

func (s *DonationService) CreateDonation(donation *models.Donation) error {
	return s.db.Create(donation).Error
}

// ListDonations retrieves all donations with donor details
func (s *DonationService) ListDonations() ([]models.Donation, error) {
	var donations []models.Donation
	if err := s.db.Preload("Donor").Find(&donations).Error; err != nil {
		return nil, err
	}
	return donations, nil
}

// UploadDonationImage uploads the donation image and returns the image URL
func (s *DonationService) UploadDonationImage(uploadDir string, c *gin.Context) (string, error) {
	return utils.UploadFile(c, uploadDir)
}

// CreateDonationRequest handles the logic for creating a donation request
func (s *DonationService) CreateDonationRequest(request *models.DonationRequest) error {
	return s.db.Create(request).Error
}

// UpdateDonationRequestStatus handles the logic for updating a donation request status
func (s *DonationService) UpdateDonationRequestStatus(id string, status string) error {
	var request models.DonationRequest
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	if err := s.db.First(&request, idInt).Error; err != nil {
		return err
	}
	request.Status = status
	return s.db.Save(&request).Error
}

// ListDonationRequests retrieves all donation requests
func (s *DonationService) ListDonationRequests() ([]models.DonationRequest, error) {
	var requests []models.DonationRequest
	if err := s.db.Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

// GetDonationRequestByID fetches a donation request by its ID and preloads the related Donation
func (s *DonationService) GetDonationRequestByID(requestID string) (*models.DonationRequest, error) {
	var request models.DonationRequest
	if err := s.db.Preload("Donation").First(&request, requestID).Error; err != nil {
		return nil, err // Return nil and the error if not found
	}
	return &request, nil // Return the found donation request
}
