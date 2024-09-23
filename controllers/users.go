package controllers

import (
	"fmt"
	"net/http"

	"github.com/elishambadi/sharebite/models"
	"github.com/elishambadi/sharebite/services"
	"github.com/elishambadi/sharebite/utils"
	"github.com/gin-gonic/gin"
)

var usersService *services.UserService

// Interface to be satisfied by any userService
type UserService interface {
	GetUsers() ([]models.User, error)
}

func GetUsers1(c *gin.Context) {
	// Gets users
	users, err := usersService.GetUsers()

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

func GetUsersHandler(service UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Gets users
		users, err := service.GetUsers()

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

}

func CreateUser(c *gin.Context) {
	// Create user
	// Link here to create user service

	err := usersService.CreateUser(c)
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
	user, err := usersService.GetUserById(userId)
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
	err := usersService.DeleteUserById(userId)
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
	token, err := usersService.AuthenticateUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{
			"message": fmt.Sprintf("Error authenticating: %s", err),
			"token":   "",
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Authenticated Successfully",
		"token":   token,
	})
}

func ResetUserPassword(ctx *gin.Context) {
	err := usersService.ResetUserPassword(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": fmt.Sprintf("Error resetting user password: %s", err),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Password reset successfully",
	})
}

func uploadUserProfile(c *gin.Context) {
	uploadDir := "./uploads/profile" // Directory to save uploaded files
	imageURL, err := utils.UploadFile(c, uploadDir)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_url": imageURL})
}

func Dashboard(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the dashboard!", "user": user})
}
