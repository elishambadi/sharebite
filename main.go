package main

import (
	"github.com/elishambadi/sharebite/config"
	"github.com/elishambadi/sharebite/db"
	"github.com/elishambadi/sharebite/routes"
	"github.com/elishambadi/sharebite/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := utils.SetupLogger()
	defer logger.Sync() // flushes any buffered logs

	r := gin.Default()

	config.LoadConfig()

	routes.SetupRoutes(r, logger)

	db.ConnectDB()

	r.Run(":8080")
}
