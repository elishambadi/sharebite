package models

import (
	"time"

	"fmt"
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

// GenerateFakeDonations creates 50 fake donations with randomization
func GenerateFakeDonations() []Donation {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Data for randomization
	foodTypes := []string{"Canned Beans", "Fresh Bread", "Rice", "Pasta", "Canned Soup", "Fresh Vegetables", "Milk", "Canned Tuna", "Cereal", "Frozen Meat"}
	locations := []string{"123 Charity St", "456 Community Rd", "789 Shelter Ln", "101 Main St", "202 Oak Ave", "303 Pine Blvd", "404 Maple Dr", "505 Birch Rd", "606 Cedar St", "707 Elm St"}
	urgencies := []string{"Low", "Medium", "High"}

	var donations []Donation

	var i uint

	for i = 1; i <= 50; i++ {
		foodType := foodTypes[rand.Intn(len(foodTypes))]                  // Random food type
		location := locations[rand.Intn(len(locations))]                  // Random location
		urgency := urgencies[rand.Intn(len(urgencies))]                   // Random urgency level
		quantity := rand.Intn(101) + 10                                   // Random quantity between 10 and 110
		expiration := time.Now().AddDate(0, rand.Intn(12), rand.Intn(30)) // Random expiration date within 1 year

		// Random donor ID (1 to 5)
		donorID := uint(rand.Intn(2) + 1)

		// Image URL based on the donation ID
		imageURL := fmt.Sprintf("https://example.com/images/donation%d.jpg", i)

		// Create a fake donation
		donation := NewDonation(
			i,          // ID
			foodType,   // FoodType
			quantity,   // Quantity
			expiration, // Expiration date
			location,   // Location
			urgency,    // Urgency
			donorID,    // DonorID
			imageURL,   // Image URL
		)

		donations = append(donations, donation)
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
