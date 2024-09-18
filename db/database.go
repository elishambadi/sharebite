package db

import (
	"fmt"
	"log"

	"github.com/elishambadi/sharebite/config"
	"github.com/elishambadi/sharebite/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := config.AppConfig.DBHost
	user := config.AppConfig.DBUser
	dbName := config.AppConfig.DBName
	password := config.AppConfig.DBPassword

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbName, password)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	database.AutoMigrate(&models.User{}) // Auto-migrate the user model

	DB = database
}
