package controllers

import (
	"fmt"
	"net/http"

	"github.com/elishambadi/sharebite/models"
	"github.com/elishambadi/sharebite/utils"
	"github.com/gin-gonic/gin"
)

// Interface to be satisfied by any userService
type UserService interface {
	GetUsers() ([]models.User, error)
	CreateUser(newUser models.User) error
	GetUserById(id string) (models.User, error)
	DeleteUserById(id string) error
	AuthenticateUser(ctx *gin.Context) (token string, error error)
	ResetUserPassword(ctx *gin.Context) error
	GetUserFromRequest(c *gin.Context) (models.User, error)
}

// GetUsersHandler godoc
// @Summary Get all users
// @Description Get the list of users
// @Tags users
// @Produce  json
// @Success 200 {object} map[string]interface{} "request successful"
// @Failure 500 {object} map[string]interface{} "error getting users"
// @Router /users [get]
func GetUsersHandler(service UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Gets users
		users, err := service.GetUsers()

		// Return response
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error getting users",
				"users":   users,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "request successful",
			"users":   users,
		})
	}

}

func CreateUserHandler(userService UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User

		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("error creating new user: %s", err),
			})
		}

		err := userService.CreateUser(newUser)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("error creating new user: %s", err),
			})
			return
		}

		//
		c.JSON(http.StatusCreated, gin.H{
			"message": "user created successfully",
		})
	}
}

func GetUserByIdHandler(userService UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("id")
		user, err := userService.GetUserById(userId)
		if err != nil {
			// record not found error
			if err.Error() == "record not found" {
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": "no user found for the given parameters",
					"user":    []models.User{},
				})
				return
			}

			// Any other error
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error fetching user",
				"user":    []models.User{},
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "user fetched successfully",
			"user":    user,
		})
	}
}

func DeleteUserByIdHandler(userService UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("id")
		err := userService.DeleteUserById(userId)
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
}

func AuthenticateUserHandler(userService UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := userService.AuthenticateUser(ctx)
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
}

func ResetUserPasswordHandler(userService UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := userService.ResetUserPassword(ctx)
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
}

func uploadUserProfilePhotoHandler(userService UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uploadDir := "./uploads/profile" // Directory to save uploaded files
		imageURL, err := utils.UploadFile(ctx, uploadDir)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"image_url": imageURL})
	}
}

func DashboardHandler(userService UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _ := ctx.Get("user")
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the dashboard!", "user": user})
	}
}
