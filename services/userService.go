package services

import (
	"errors"
	"log"
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
