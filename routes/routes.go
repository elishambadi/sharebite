package routes

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/elishambadi/sharebite/controllers"
	"github.com/elishambadi/sharebite/middlewares"
	"github.com/elishambadi/sharebite/services"
	templates "github.com/elishambadi/sharebite/templates/client"
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

func SetupRoutes(r *gin.Engine, logger *zap.Logger, userController controllers.UserController, userService services.UserService, donationService services.DonationService) {
	// Add logger to all routes
	r.Use(middlewares.LoggerMiddleware(logger))

	r.Static("/assets", "./assets")

	// Root route
	r.GET("/", func(ctx *gin.Context) {
		render(ctx, http.StatusOK, templates.Base("Dashboard", templates.Dashboard()))
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
		apiRoutes.POST("/donations", controllers.CreateDonationHandler(&donationService, userService))
		apiRoutes.POST("/upload-donation-image", controllers.UploadDonationImageHandler(&donationService))
		apiRoutes.POST("/donation-requests", controllers.CreateDonationRequestHandler(&donationService, userService))
		apiRoutes.PUT("/donation-requests/:id/status", controllers.UpdateDonationRequestStatusHandler(&donationService, userService))
		apiRoutes.GET("/donation-requests", controllers.ListDonationRequestsHandler(&donationService))

		apiRoutes.POST("/signup", userController.CreateUserHandler())
		apiRoutes.POST("/login", userController.AuthenticateUserHandler())
		apiRoutes.GET("/donations", controllers.ListDonationsHandler(&donationService))
		apiRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// App Routes (Frontend-specific)
	appRoutes := r.Group("/app")
	appRoutes.Use()
	{
		appRoutes.GET("/signup", func(c *gin.Context) {
			render(c, http.StatusOK, templates.Base("SignUp", templates.SignUp()))
		})

		appRoutes.GET("/login", func(c *gin.Context) {
			render(c, http.StatusOK, templates.Base("Login", templates.Login()))
		})

		// Dashboard
		appRoutes.GET("/dashboard", func(c *gin.Context) {
			render(c, http.StatusOK, templates.Base("Dashboard", templates.Dashboard()))
		})

		// Create a new donation
		appRoutes.GET("/donations/new", func(c *gin.Context) {
			render(c, http.StatusOK, templates.Base("Give a Donation", templates.AddDonation()))
		})

		// Create a new donation
		appRoutes.GET("/donations", func(c *gin.Context) {
			donations, _ := donationService.ListDonations()
			render(c, http.StatusOK, templates.Base("Donations", templates.DonationsList(donations)))
		})

		// View donation details
		appRoutes.GET("/donations/:id", func(c *gin.Context) {
			donationID := c.Param("id")
			donation, _ := donationService.GetDonationByID(donationID)

			render(c, http.StatusOK, templates.Base("View Donation", templates.DonationDetail(*donation)))
		})

		// List donation requests
		appRoutes.GET("/donation-requests", func(c *gin.Context) {
			render(c, http.StatusOK, templates.Base("Donation Requests", templates.DonationRequests()))
		})

		// User profile
		appRoutes.GET("/profile", func(c *gin.Context) {
			render(c, http.StatusOK, templates.Base("My profile", templates.Profile()))
		})
	}
}
