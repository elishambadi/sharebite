package main

import (
	"github.com/elishambadi/sharebite/config"
	"github.com/elishambadi/sharebite/controllers"
	"github.com/elishambadi/sharebite/db"
	repository "github.com/elishambadi/sharebite/repositories"
	"github.com/elishambadi/sharebite/routes"
	"github.com/elishambadi/sharebite/services"
	"github.com/elishambadi/sharebite/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := utils.SetupLogger()
	defer logger.Sync() // flushes any buffered logs

	r := gin.Default()

	config.LoadConfig()

	db.ConnectDB()

	userRepo := repository.NewGormUserRepository(db.DB, logger)
	userService := services.NewUserService(userRepo, logger)
	userController := controllers.NewUserController(*userService)

	routes.SetupRoutes(r, logger, userController)

	r.Run(":8080")
}
