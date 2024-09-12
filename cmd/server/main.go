package main

import (
	"log"

	"github.com/elishambadi/sharebite/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set up API routes
	api.SetupRoutes(router)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
