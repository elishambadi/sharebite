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
func NewUser(id uint, name, email, password, userType string) User {
	return User{
		Model:     gorm.Model{ID: id},
		Name:      name,
		Email:     email,
		Password:  password, // Optionally hash the password here
		Type:      userType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		APIToken:  uuid.New().String(), // Generate a random token, or use some static value
	}
}

// Creates two fake users for testing
//
// User 1: User(1, "Alice", "alice@example.com", 12345678, "DONOR")
//
// User 2: User(2, "Bob", "bob@example.com", 12345678, "RECIPIENT")
func FakeUsers() []User {
	hashedPassword, _ := utils.HashPassword("password")
	return []User{
		NewUser(1, "Alice", "alice@example.com", hashedPassword, "DONOR"),
		NewUser(2, "Bob", "bob@example.com", hashedPassword, "RECIPIENT"),
	}
}
