package services

import (
	"log"
	"time"

	"github.com/elishambadi/sharebite/db"
	"github.com/elishambadi/sharebite/models"
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
	newUser := models.User{
		Name:      "Elisha Mbadi",
		Email:     "embadi@gmail.com",
		Password:  "43242342342",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := db.DB.Create(&newUser)
	if result.Error != nil {
		log.Println("Error creating user: ", result.Error)
		return result.Error
	} else {
		log.Print("User created successfully.")
		return nil
	}
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
