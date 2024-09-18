package main

import (
	"github.com/elishambadi/sharebite/config"
	"github.com/elishambadi/sharebite/db"
	"github.com/elishambadi/sharebite/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.LoadConfig()

	routes.SetupRoutes(r)

	db.ConnectDB()

	r.Run(":8080")
}
