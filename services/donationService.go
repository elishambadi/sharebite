package services

import (
	"github.com/elishambadi/sharebite/models" // Adjust the import path as needed
	repository "github.com/elishambadi/sharebite/repositories"
	"github.com/elishambadi/sharebite/utils"
	"github.com/gin-gonic/gin"
)

type DonationService struct {
	repo repository.DonationRepository
}

func NewDonationService(repo repository.DonationRepository) *DonationService {
	return &DonationService{repo: repo}
}

func (s *DonationService) CreateDonation(donation *models.Donation) error {
	return s.repo.CreateDonation(donation)
}

// ListDonations retrieves all donations with donor details
func (s *DonationService) ListDonations() ([]models.Donation, error) {
	return s.repo.FindAll()
}

// UploadDonationImage uploads the donation image and returns the image URL
func (s *DonationService) UploadDonationImage(uploadDir string, c *gin.Context) (string, error) {
	return utils.UploadFile(c, uploadDir)
}

// CreateDonationRequest handles the logic for creating a donation request
func (s *DonationService) CreateDonationRequest(request *models.DonationRequest) error {
	return s.repo.CreateDonationRequest(request)
}

// UpdateDonationRequestStatus handles the logic for updating a donation request status
func (s *DonationService) UpdateDonationRequestStatus(id string, status string) error {
	return s.repo.UpdateDonationRequestStatus(id, status)
}

// ListDonationRequests retrieves all donation requests
func (s *DonationService) ListDonationRequests() ([]models.DonationRequest, error) {
	return s.repo.ListDonationRequests()
}

// GetDonationRequestByID fetches a donation request by its ID and preloads the related Donation
func (s *DonationService) GetDonationRequestByID(requestID string) (*models.DonationRequest, error) {
	return s.repo.GetDonationRequestByID(requestID)
}
