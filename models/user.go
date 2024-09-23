package models

import (
	"errors"
	"time"

	"github.com/elishambadi/sharebite/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique; not null"`
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
		return errors.New("invalid type: must either be 'DONOR' or 'RECIPIENT'")
	}
	return nil
}

// Factory method to return a new user
func NewUser(name, email, password, userType string) User {
	return User{
		Name:      name,
		Email:     email,
		Password:  password, // Optionally hash the password here
		Type:      userType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		APIToken:  uuid.New().String(), // Generate a random token, or use some static value
	}
}

func FakeUsers() []User {
	hashedPassword, _ := utils.HashPassword("password")
	return []User{
		NewUser("Alice", "alice@example.com", hashedPassword, "DONOR"),
		NewUser("Bob", "bob@example.com", hashedPassword, "RECIPIENT"),
	}
}
