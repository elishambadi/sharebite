package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	GinMode    string
	SeedDB     bool
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env: %s", err)
	}

	AppConfig = &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		GinMode:    os.Getenv("GIN_MODE"),
		SeedDB:     getSeedDBValue(os.Getenv("SEED_DB")),
	}

	gin.SetMode(AppConfig.GinMode)
}

func getSeedDBValue(seedDBString string) bool {
	return seedDBString == "true"
}
