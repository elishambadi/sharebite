package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `gorm:"unique;not null"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
