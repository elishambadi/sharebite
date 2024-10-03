package services

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/elishambadi/sharebite/db"
	"github.com/elishambadi/sharebite/models"
	repository "github.com/elishambadi/sharebite/repositories"
	"github.com/elishambadi/sharebite/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserService struct {
	repo   *repository.GormUserRepository
	logger *zap.Logger
}

func NewUserService(repo *repository.GormUserRepository, logger *zap.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (u *UserService) GetUsers() ([]models.User, error) {
	return u.repo.FindAll()
}

func (u *UserService) CreateUser(newUser models.User) error {
	hashedPw, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = hashedPw

	return u.repo.Create(newUser)
}

func (u *UserService) AuthenticateUser(ctx *gin.Context) (string, error) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		return "", err
	}

	foundUser, err := u.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	passwordValid := utils.CheckPassword(foundUser.Password, user.Password)
	if !passwordValid {
		return "", errors.New("invalid password")
	}

	userIdString := strconv.Itoa(int(foundUser.ID))

	token, err := utils.CreateJWT(userIdString, []string{"admin"})
	if err != nil {
		return "", err
	}

	if dbError := u.repo.UpdateAPIToken(&foundUser, token); dbError != nil {
		return "", dbError
	}

	ctx.Set("user", foundUser)

	return token, nil
}

func (u *UserService) ResetUserPassword(c *gin.Context) error {
	var userDetails models.User

	if err := c.ShouldBindJSON(&userDetails); err != nil {
		return err
	}

	u.logger.Info("Resetting user password", zap.String("userEmail", userDetails.Email))

	user, err := u.GetUserByEmail(userDetails.Email)
	if err != nil {
		return err
	}

	newPassword, err := utils.HashPassword(userDetails.Password)
	if err != nil {
		return err
	}

	user.Password = newPassword
	if err := db.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

// Gets a user from every authenticated request
func (u *UserService) GetUserFromRequest(c *gin.Context) (models.User, error) {
	var userModel models.User

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "",
		})
		return userModel, errors.New("unauthenticated. Please authenticate your request")
	}

	user, ok := user.(models.User)
	if !ok {
		return userModel, errors.New("error retrieving user data")
	}

	return user.(models.User), nil
}

func (u *UserService) GetUserById(id string) (models.User, error) {
	return u.repo.GetUserById(id)
}

func (u *UserService) GetUserByEmail(email string) (models.User, error) {
	return u.repo.GetUserByEmail(email)
}

func (u *UserService) DeleteUserById(id string) error {
	return u.repo.DeleteUserById(id)
}
