package services

import "github.com/elishambadi/sharebite/models"

func GetUser() ([]models.User, error) {
	// Gets a user from the DB
	users := []models.User{}

	user := models.User{
		Id:       1,
		Name:     "Mark Lewis",
		Email:    "marklewis@gmail.com",
		Password: "12345678",
	}
	users = append(users, user)

	return users, nil
}
