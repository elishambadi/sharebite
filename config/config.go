package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	GinMode    string
	SeedDB     bool
	AppURL     string
}

var AppConfig *Config

func LoadConfig() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env: %s", err)
	// }

	AppConfig = &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		GinMode:    os.Getenv("GIN_MODE"),
		SeedDB:     getSeedDBValue(os.Getenv("SEED_DB")),
		AppURL:     os.Getenv("APP_URL"),
	}

	gin.SetMode(AppConfig.GinMode)
}

func getSeedDBValue(seedDBString string) bool {
	return seedDBString == "true"
}
