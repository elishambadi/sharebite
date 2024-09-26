package models

import (
	"time"

	"math/rand"

	"gorm.io/gorm"
)

// Donation struct represents a food donation record
type Donation struct {
	gorm.Model
	FoodType   string    `json:"food_type" binding:"required"`  // Type of food being donated
	Quantity   int       `json:"quantity" binding:"required"`   // Quantity of food
	Expiration time.Time `json:"expiration" binding:"required"` // Expiration date of the food
	Location   string    `json:"location" binding:"required"`   // Location of the donation
	Urgency    string    `json:"urgency"`                       // Urgency of the donation (optional, can use Low, Medium, High)
	DonorID    uint      `json:"donor_id"`                      // Foreign key to associate with User
	Donor      User      `gorm:"foreignKey:DonorID"`            // The donor, relation to User model
	CreatedAt  time.Time `json:"created_at"`                    // Timestamp for creation
	UpdatedAt  time.Time `json:"updated_at"`                    // Timestamp for update
	ImageURL   string    `json:"image_url"`
}

type DonationRequest struct {
	gorm.Model
	DonationID  uint      `json:"donation_id"`  // Foreign key to Donation
	RecipientID uint      `json:"recipient_id"` // Foreign key to User or Recipient
	Status      string    `json:"status"`       // Status of the request (e.g., Pending, Approved, Rejected)
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Donation    Donation
}

func NewDonation(id uint, foodType string, quantity int, expiration time.Time, location string, urgency string, donorID uint, imageURL string) Donation {
	return Donation{
		Model:      gorm.Model{ID: id},
		FoodType:   foodType,
		Quantity:   quantity,
		Expiration: expiration,
		Location:   location,
		Urgency:    urgency,
		DonorID:    donorID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		ImageURL:   imageURL,
	}
}

// NewDonationRequest creates a new DonationRequest instance
func NewDonationRequest(id uint, donationID uint, recipientID uint, status string) DonationRequest {
	return DonationRequest{
		Model:       gorm.Model{ID: id},
		DonationID:  donationID,
		RecipientID: recipientID,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// GenerateFakeDonations creates 3 fake donations using NewDonation
func GenerateFakeDonations() []Donation {
	donations := []Donation{
		NewDonation(1, "Canned Beans", 50, time.Now().AddDate(0, 3, 0), "123 Charity St", "High", 1, "https://example.com/images/donation1.jpg"),
		NewDonation(2, "Fresh Bread", 20, time.Now().AddDate(0, 0, 3), "456 Community Rd", "Medium", 2, "https://example.com/images/donation2.jpg"),
		NewDonation(3, "Rice", 100, time.Now().AddDate(0, 6, 0), "789 Shelter Ln", "Low", 1, "https://example.com/images/donation3.jpg"),
	}

	return donations
}

// GenerateFakeDonationRequests creates 3 fake donation requests using NewDonationRequest
func GenerateFakeDonationRequests() []DonationRequest {
	requests := []DonationRequest{
		NewDonationRequest(1, 1, uint(rand.Intn(100)+3), "Pending"),
		NewDonationRequest(2, 2, uint(rand.Intn(100)+3), "Approved"),
		NewDonationRequest(3, 3, uint(rand.Intn(100)+3), "Rejected"),
	}

	return requests
}
