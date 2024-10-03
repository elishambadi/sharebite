package controllers

import (
	"fmt"
	"net/http"

	"github.com/elishambadi/sharebite/models"
	"github.com/elishambadi/sharebite/services"
	"github.com/elishambadi/sharebite/utils"
	"github.com/gin-gonic/gin"
)

// UserController interface
type UserController interface {
	GetUsersHandler() gin.HandlerFunc
	CreateUserHandler() gin.HandlerFunc
	GetUserByIdHandler() gin.HandlerFunc
	DeleteUserByIdHandler() gin.HandlerFunc
	AuthenticateUserHandler() gin.HandlerFunc
	ResetUserPasswordHandler() gin.HandlerFunc
	uploadUserProfilePhotoHandler() gin.HandlerFunc
	DashboardHandler() gin.HandlerFunc
}

// userController is the concrete implementation of UserController.
type userController struct {
	userService services.UserService
}

// NewUserController creates a new UserController instance.
func NewUserController(userService services.UserService) UserController {
	return &userController{userService: userService}
}

// GetUsersHandler godoc
// @Summary Get all users
// @Description Get the list of users
// @Tags users
// @Produce  json
// @Success 200 {object} map[string]interface{} "request successful"
// @Failure 500 {object} map[string]interface{} "error getting users"
// @Router /users [get]
func (c *userController) GetUsersHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Gets users
		users, err := c.userService.GetUsers()

		// Return response
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error getting users",
				"users":   users,
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "request successful",
			"users":   users,
		})
	}

}

func (c *userController) CreateUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newUser models.User

		if err := ctx.ShouldBindJSON(&newUser); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("error creating new user: %s", err),
			})
		}

		err := c.userService.CreateUser(newUser)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("error creating new user: %s", err),
			})
			return
		}

		//
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "user created successfully",
		})
	}
}

func (c *userController) GetUserByIdHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("id")
		user, err := c.userService.GetUserById(userId)
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

func (c *userController) DeleteUserByIdHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("id")
		err := c.userService.DeleteUserById(userId)
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

func (c *userController) AuthenticateUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := c.userService.AuthenticateUser(ctx)
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

func (c *userController) ResetUserPasswordHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := c.userService.ResetUserPassword(ctx)
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

func (c *userController) uploadUserProfilePhotoHandler() gin.HandlerFunc {
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

func (c *userController) DashboardHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _ := ctx.Get("user")
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the dashboard!", "user": user})
	}
}
