package routes

import (
	"github.com/elishambadi/sharebite/controllers"
	"github.com/elishambadi/sharebite/middlewares"
	"github.com/elishambadi/sharebite/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// use anonymous func to return
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to ShareBite API!",
		})
	})

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsersHandler(&services.UserService{}))
		userRoutes.GET("/:id", controllers.GetUserByIdHandler(&services.UserService{}))
		userRoutes.DELETE("/:id", controllers.DeleteUserByIdHandler(&services.UserService{}))
		userRoutes.POST("/reset-password", controllers.ResetUserPasswordHandler(&services.UserService{}))
	}

	protectedRoutes := r.Group("/app")
	protectedRoutes.Use(middlewares.CheckUserRole)
	{
		protectedRoutes.GET("/dashboard", controllers.DashboardHandler(&services.UserService{}))
		protectedRoutes.POST("/donations", controllers.CreateDonation)
		protectedRoutes.POST("/upload-donation-image", controllers.UploadDonationImage)
		protectedRoutes.POST("/donation-requests", controllers.CreateDonationRequest)
		protectedRoutes.PUT("/donation-requests/:id/status", controllers.UpdateDonationRequestStatus)
		protectedRoutes.GET("/donation-requests", controllers.ListDonationRequests)
	}

	r.POST("/signup", controllers.CreateUserHandler(&services.UserService{}))
	r.POST("/login", controllers.AuthenticateUserHandler(&services.UserService{}))
	r.GET("/donations", controllers.ListDonations)
}
