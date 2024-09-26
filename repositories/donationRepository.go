package repository

import (
	"strconv"

	"github.com/elishambadi/sharebite/models"
	"gorm.io/gorm"
)

type DonationRepository struct {
	db *gorm.DB
}

// constructor fn
func NewDonationRepository(db *gorm.DB) *DonationRepository {
	return &DonationRepository{
		db: db,
	}
}

func (dr *DonationRepository) FindAll() ([]models.Donation, error) {
	var donations []models.Donation

	// Preload the relationship
	if err := dr.db.Preload("Donor").Find(&donations).Error; err != nil {
		return []models.Donation{}, err
	}

	return donations, nil
}

func (dr *DonationRepository) CreateDonation(donation *models.Donation) error {
	return dr.db.Create(&donation).Error
}

func (dr *DonationRepository) CreateDonationRequest(request *models.DonationRequest) error {
	return dr.db.Create(&request).Error
}

func (dr *DonationRepository) UpdateDonationRequestStatus(id string, status string) error {
	var request models.DonationRequest
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	if err := dr.db.First(&request, idInt).Error; err != nil {
		return err
	}
	request.Status = status
	return dr.db.Save(&request).Error
}

func (dr *DonationRepository) ListDonationRequests() ([]models.DonationRequest, error) {
	var requests []models.DonationRequest
	if err := dr.db.Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (dr *DonationRepository) GetDonationRequestByID(requestID string) (*models.DonationRequest, error) {
	var request models.DonationRequest
	if err := dr.db.Preload("Donation").First(&request, requestID).Error; err != nil {
		return nil, err
	}
	return &request, nil
}
