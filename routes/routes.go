package routes

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/elishambadi/sharebite/controllers"
	"github.com/elishambadi/sharebite/middlewares"
	"github.com/elishambadi/sharebite/services"
	"github.com/elishambadi/sharebite/templates"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// renders templ templates
func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func SetupRoutes(r *gin.Engine, logger *zap.Logger, userController controllers.UserController, userService services.UserService) {
	// Add logger to all routes
	r.Use(middlewares.LoggerMiddleware(logger))

	r.Static("/assets", "./assets")

	// Root route
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to ShareBite API!",
		})
	})

	// API Routes
	apiRoutes := r.Group("/api")
	{
		userRoutes := apiRoutes.Group("/users")
		{
			userRoutes.GET("/", userController.GetUsersHandler())
			userRoutes.GET("/:id", userController.GetUserByIdHandler())
			userRoutes.DELETE("/:id", userController.DeleteUserByIdHandler())
			userRoutes.POST("/reset-password", userController.ResetUserPasswordHandler())
		}

		apiRoutes.GET("/dashboard", userController.DashboardHandler())
		apiRoutes.POST("/donations", controllers.CreateDonationHandler(&services.DonationService{}, userService))
		apiRoutes.POST("/upload-donation-image", controllers.UploadDonationImageHandler(&services.DonationService{}))
		apiRoutes.POST("/donation-requests", controllers.CreateDonationRequestHandler(&services.DonationService{}, userService))
		apiRoutes.PUT("/donation-requests/:id/status", controllers.UpdateDonationRequestStatusHandler(&services.DonationService{}, userService))
		apiRoutes.GET("/donation-requests", controllers.ListDonationRequestsHandler(&services.DonationService{}))

		apiRoutes.POST("/signup", userController.CreateUserHandler())
		apiRoutes.POST("/login", userController.AuthenticateUserHandler())
		apiRoutes.GET("/donations", controllers.ListDonationsHandler(&services.DonationService{}))
		apiRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// App Routes (Frontend-specific)
	appRoutes := r.Group("/app")
	appRoutes.Use(middlewares.CheckUserRole)
	{
		// Dashboard
		appRoutes.GET("/dashboard", func(c *gin.Context) {
			render(c, http.StatusOK, templates.Base(templates.Dashboard()))
		})

		// Create a new donation
		appRoutes.GET("/donations/new", func(c *gin.Context) {
			render(c, http.StatusOK, templates.Base(templates.NewDonation()))
		})

		// View donation details
		appRoutes.GET("/donations/:id", func(c *gin.Context) {
			donationID := c.Param("id")
			render(c, http.StatusOK, templates.Base(templates.DonationDetail(donationID)))
		})

		// List donation requests
		appRoutes.GET("/donation-requests", func(c *gin.Context) {
			render(c, http.StatusOK, templates.Base(templates.DonationRequests()))
		})

		// User profile
		appRoutes.GET("/profile", func(c *gin.Context) {
			render(c, http.StatusOK, templates.Base(templates.Profile()))
		})
	}
}
