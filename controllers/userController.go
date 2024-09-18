package controllers

import (
	"net/http"

	"github.com/elishambadi/sharebite/services"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	// Gets users
	users, err := services.GetUser()

	// Return response
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting Users",
			"users":   users,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Request successful",
		"users":   users,
	})
}

func CreateUser(c *gin.Context) {
	// Create user
	// Link here to create user service

	//
	c.JSON(http.StatusCreated, gin.H{
		"message": "Request successful",
		"user":    []string{"Jack"},
	})
}
