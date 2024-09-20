package routes

import (
	"github.com/elishambadi/sharebite/controllers"
	"github.com/elishambadi/sharebite/middlewares"
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
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUserById)
		userRoutes.DELETE("/:id", controllers.DeleteUserById)
	}

	protectedRoutes := r.Group("/app")
	protectedRoutes.Use(middlewares.CheckUserRole)
	{
		protectedRoutes.GET("/dashboard", controllers.Dashboard)
		r.POST("/donations", controllers.CreateDonation)
	}

	r.POST("/signup", controllers.CreateUser)
	r.POST("/login", controllers.AuthenticateUser)
	r.GET("/donations", controllers.ListDonations)
}
