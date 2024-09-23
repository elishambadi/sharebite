package repository

import (
	"github.com/elishambadi/sharebite/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) FindAll() ([]models.User, error) {
	var users []models.User

	result := r.db.Find(&users)
	if result.Error != nil {
		return []models.User{}, result.Error
	}
	return users, nil
}
