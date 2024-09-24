package repository

import (
	"log"

	"github.com/elishambadi/sharebite/db"
	"github.com/elishambadi/sharebite/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	Create(newUser models.User) error
	UpdateAPIToken(user models.User, ApiToken string) error
	GetUserById(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	DeleteUserById(id string) error
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

func (r *GormUserRepository) Create(newUser models.User) error {
	result := r.db.Create(&newUser)
	if result.Error != nil {
		log.Println("Error creating user: ", result.Error)
		return result.Error
	} else {
		log.Print("User created successfully.")
		return nil
	}
}

func (r *GormUserRepository) UpdateAPIToken(user models.User, ApiToken string) error {
	user.APIToken = ApiToken
	if err := r.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormUserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("Error getting using with ID %s. Error: %s.\n", email, err)
		return models.User{}, err
	}

	log.Printf("User fetched successfully ID %s.\n", email)
	return user, nil
}

func (r *GormUserRepository) GetUserById(id string) (models.User, error) {
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		log.Printf("Error getting using with ID %s. Error: %s.\n", id, result.Error)
		return models.User{}, result.Error
	}

	log.Printf("User fetched successfully ID %s.\n", id)
	return user, nil
}

func (r *GormUserRepository) DeleteUserById(id string) error {
	var user models.User
	user, err := r.GetUserById(id)
	if err != nil {
		return err
	}

	delResult := db.DB.Delete(&user)
	if delResult.Error != nil {
		log.Printf("Error deleting user with ID %s. Error: %s.\n", id, delResult.Error)
		return delResult.Error
	} else {
		log.Printf("User deleted successfully ID %s.\n", id)
	}

	return nil
}
