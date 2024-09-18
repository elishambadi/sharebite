package controllers

import (
	"net/http"

	"github.com/elishambadi/sharebite/services"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	// Gets users
	users, err := services.GetUsers()

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

	err := services.CreateUser(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Error creating new user",
		})
	}

	//
	c.JSON(http.StatusCreated, gin.H{
		"message": "User Created Successfully",
	})
}

func GetUserById(ctx *gin.Context) {
	userId := ctx.Param("id")
	user, err := services.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching user",
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "User fetched successfully",
		"user":    user,
	})
}
