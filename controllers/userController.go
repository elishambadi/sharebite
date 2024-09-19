package controllers

import (
	"fmt"
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
			"message": fmt.Sprintf("Error creating new user: %s", err),
		})
		return
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
		// record not found error
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusAccepted, gin.H{
				"message": "No user found for the given parameters",
			})
			return
		}

		// Any other error
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

func DeleteUserById(ctx *gin.Context) {
	userId := ctx.Param("id")
	err := services.DeleteUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error deleting user: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "User deleted successfully",
	})
}

func AuthenticateUser(ctx *gin.Context) {
	token, err := services.AuthenticateUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{
			"message": fmt.Sprintf("Error authenticating: %s", err),
			"token":   "",
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Authenticated Successfully",
		"token":   token,
	})
}

func Dashboard(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the dashboard!", "user": user})
}
