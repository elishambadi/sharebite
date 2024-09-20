package models

import (
	"time"
)

// Donation struct represents a food donation record
type Donation struct {
	ID         uint      `gorm:"primaryKey"`
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
	ID          uint      `gorm:"primaryKey"`
	DonationID  uint      `json:"donation_id"`  // Foreign key to Donation
	RecipientID uint      `json:"recipient_id"` // Foreign key to User or Recipient
	Status      string    `json:"status"`       // Status of the request (e.g., Pending, Approved, Rejected)
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Donation    Donation
}
