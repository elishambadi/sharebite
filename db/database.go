package db

import (
	"fmt"
	"log"
	"time"

	"github.com/elishambadi/sharebite/config"
	"github.com/elishambadi/sharebite/models"
	"github.com/elishambadi/sharebite/utils"

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

	database.AutoMigrate(&models.User{}, &models.Donation{}, &models.DonationRequest{}) // Auto-migrate the user model

	DB = database

	if config.AppConfig.SeedDB {
		seedDB()
	}
}

func seedDB() {
	hashedPassword, _ := utils.HashPassword("password")
	users := []models.User{
		{Name: "Alice", Email: "alice@example.com", Password: hashedPassword, Type: "DONOR", APIToken: "token1"},
		{Name: "Bob", Email: "bob@example.com", Password: hashedPassword, Type: "RECIPIENT", APIToken: "token2"},
	}

	for idx, user := range users {
		if err := DB.Create(&user).Error; err != nil {
			log.Fatalf("failed to create user: %v", err)
		}
		users[idx] = user
	}

	// Seed donations
	donations := []models.Donation{
		{
			FoodType:   "Vegetables",
			Quantity:   10,
			Expiration: time.Now().AddDate(0, 1, 0),
			Location:   "Warehouse A",
			Urgency:    "Low",
			DonorID:    users[0].ID, // Use Alice's ID
			ImageURL:   "http://example.com/image1.jpg",
		},
		{
			FoodType:   "Canned Beans",
			Quantity:   20,
			Expiration: time.Now().AddDate(0, 0, 15),
			Location:   "Warehouse B",
			Urgency:    "High",
			DonorID:    users[0].ID, // Use Alice's ID
			ImageURL:   "http://example.com/image2.jpg",
		},
	}

	for idx, donation := range donations {
		if err := DB.Create(&donation).Error; err != nil {
			log.Fatalf("failed to create donation: %v", err)
		}

		donations[idx] = donation
	}

	// Seed donation requests
	donationRequests := []models.DonationRequest{
		{
			DonationID:  donations[0].ID, // Link to first donation
			RecipientID: users[1].ID,     // Use Bob's ID
			Status:      "PENDING",
		},
		{
			DonationID:  donations[1].ID, // Link to second donation
			RecipientID: users[1].ID,     // Use Bob's ID
			Status:      "APPROVED",
		},
	}

	for _, request := range donationRequests {
		if err := DB.Create(&request).Error; err != nil {
			log.Fatalf("failed to create donation request: %v", err)
		}
	}

	log.Println("Database seeded successfully!")
}
