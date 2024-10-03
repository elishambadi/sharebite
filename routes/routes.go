package routes

import (
	"github.com/elishambadi/sharebite/controllers"
	"github.com/elishambadi/sharebite/middlewares"
	"github.com/elishambadi/sharebite/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func SetupRoutes(r *gin.Engine, logger *zap.Logger, userController controllers.UserController) {
	// Add logger to all routes
	r.Use(middlewares.LoggerMiddleware(logger))

	// use anonymous func to return
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to ShareBite API!",
		})
	})

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", userController.GetUsersHandler())
		userRoutes.GET("/:id", userController.GetUserByIdHandler())
		userRoutes.DELETE("/:id", userController.DeleteUserByIdHandler())
		userRoutes.POST("/reset-password", userController.ResetUserPasswordHandler())
	}

	protectedRoutes := r.Group("/app")
	protectedRoutes.Use(middlewares.CheckUserRole)
	{
		protectedRoutes.GET("/dashboard", userController.DashboardHandler())
		protectedRoutes.POST("/donations", controllers.CreateDonationHandler(&services.DonationService{}, &services.UserService{}))
		protectedRoutes.POST("/upload-donation-image", controllers.UploadDonationImageHandler(&services.DonationService{}))
		protectedRoutes.POST("/donation-requests", controllers.CreateDonationRequestHandler(&services.DonationService{}, &services.UserService{}))
		protectedRoutes.PUT("/donation-requests/:id/status", controllers.UpdateDonationRequestStatusHandler(&services.DonationService{}, &services.UserService{}))
		protectedRoutes.GET("/donation-requests", controllers.ListDonationRequestsHandler(&services.DonationService{}))
	}

	r.POST("/signup", userController.CreateUserHandler())
	r.POST("/login", userController.AuthenticateUserHandler())
	r.GET("/donations", controllers.ListDonationsHandler(&services.DonationService{}))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
