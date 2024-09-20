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
}
