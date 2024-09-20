package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Type      string `json:"type"`
	CreatedAt time.Time
	UpdatedAt time.Time
	APIToken  string
	Donations []Donation `gorm:"foreignKey:DonorID"` // One-to-Many relationship with donations
}

// Performs this check before a user is saved
func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Type != "DONOR" && u.Type != "RECIPIENT" {
		return errors.New("Invalid type: must either be 'DONOR' or 'RECIPIENT'")
	}
	return nil
}
