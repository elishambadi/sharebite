package utils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupLogger() *zap.Logger {
	var logger *zap.Logger
	var err error

	if gin.Mode() == gin.ReleaseMode {
		// Writes error as JSON
		logger, err = zap.NewProduction()
	} else {
		// Writes error as readable text
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic("failed to initialize logger")
	}

	return logger
}
