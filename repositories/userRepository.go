package repository

import (
	"strconv"

	"github.com/elishambadi/sharebite/models"
	"go.uber.org/zap"
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
	db     *gorm.DB
	logger *zap.Logger
}

func NewGormUserRepository(db *gorm.DB, logger *zap.Logger) *GormUserRepository {
	return &GormUserRepository{
		db:     db,
		logger: logger,
	}
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
		r.logger.Error("Error creating user: ", zap.String("userEmail", newUser.Email), zap.Error(result.Error))
		return result.Error
	} else {
		r.logger.Info("User created successfully", zap.Uint("userId", newUser.ID))
		return nil
	}
}

func (r *GormUserRepository) UpdateAPIToken(user *models.User, ApiToken string) error {
	user.APIToken = ApiToken
	if err := r.db.Save(&user).Error; err != nil {
		r.logger.Error("Error updating API key", zap.String("userEmail", user.Email))
		return err
	}
	return nil
}

func (r *GormUserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		r.logger.Error("Error getting user by email", zap.String("userEmail", user.Email), zap.Error(err))
		return models.User{}, err
	}

	r.logger.Info("User fetched via email", zap.String("userEmail", user.Email))
	return user, nil
}

func (r *GormUserRepository) GetUserById(id string) (models.User, error) {
	var user models.User
	idUint, err := strconv.Atoi(id)
	if err != nil {
		r.logger.Error("Invalid user ID format", zap.String("userId", id), zap.Error(err))
		return models.User{}, err // return error if ID is not valid
	}

	result := r.db.Take(&user, idUint)
	if result.Error != nil {
		r.logger.Error("Error getting user by ID", zap.String("userId", id), zap.Error(result.Error))
		return models.User{}, result.Error
	}

	r.logger.Info("User fetched via ID", zap.Uint("userId", user.ID))
	return user, nil
}

func (r *GormUserRepository) DeleteUserById(id string) error {
	var user models.User
	user, err := r.GetUserById(id)
	if err != nil {
		return err
	}

	delResult := r.db.Delete(&user)
	if delResult.Error != nil {
		r.logger.Error("Error deleting user by Id", zap.String("userId", id), zap.Error(delResult.Error))
		return delResult.Error
	}

	r.logger.Info("User deleted successfully", zap.String("userId", id), zap.String("userEmail", user.Email))
	return nil
}
