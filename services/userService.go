package services

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/elishambadi/sharebite/db"
	"github.com/elishambadi/sharebite/models"
	"github.com/elishambadi/sharebite/utils"
	"github.com/gin-gonic/gin"
)

func GetUsers() ([]models.User, error) {
	// Gets a user from the DB
	var users []models.User

	result := db.DB.Find(&users)
	if result.Error != nil {
		log.Println("Error fetching users in userService: ", result.Error)
		return []models.User{}, result.Error
	} else {
		log.Println("Users fetched successfully")
	}

	return users, nil
}

func CreateUser(ctx *gin.Context) error {
	var newUser models.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		return err
	}

	hashedPw, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = hashedPw

	result := db.DB.Create(&newUser)
	if result.Error != nil {
		log.Println("Error creating user: ", result.Error)
		return result.Error
	} else {
		log.Print("User created successfully.")
		return nil
	}
}

func AuthenticateUser(ctx *gin.Context) (string, error) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		return "", err
	}

	foundUser, err := GetUserByEmail(user.Email)
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

	foundUser.APIToken = token
	db.DB.Save(&foundUser)

	ctx.Set("user", foundUser)

	return token, nil
}

func ResetUserPassword(c *gin.Context) error {
	var userDetails models.User

	if err := c.ShouldBindJSON(&userDetails); err != nil {
		return err
	}
	log.Printf("Resetting password for %s to %s.", userDetails.Email, userDetails.Password)

	user, err := GetUserByEmail(userDetails.Email)
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
func GetUserFromRequest(c *gin.Context) (models.User, error) {
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

func GetUserById(id string) (models.User, error) {
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		log.Printf("Error getting using with ID %s. Error: %s.\n", id, result.Error)
		return models.User{}, result.Error
	} else {
		log.Printf("User fetched successfully ID %s.\n", id)
	}

	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("Error getting using with ID %s. Error: %s.\n", email, err)
		return models.User{}, err
	} else {
		log.Printf("User fetched successfully ID %s.\n", email)
	}

	return user, nil
}

func DeleteUserById(id string) error {
	var user models.User
	user, err := GetUserById(id)
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
