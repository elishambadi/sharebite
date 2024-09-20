package controllers

import (
	"net/http"

	"github.com/elishambadi/sharebite/db"
	"github.com/elishambadi/sharebite/models"

	"github.com/gin-gonic/gin"
)

// CreateDonation handles POST requests to log a new food donation
func CreateDonation(c *gin.Context) {
	user, _ := c.Get("user")
	var donation models.Donation
	if err := c.ShouldBindJSON(&donation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add DonorID manually or via authentication (assume DonorID = 1 for this example)
	donation.DonorID = user.Id // Normally, you would extract this from the authenticated user context.

	// Save the donation to the database
	if err := db.DB.Create(&donation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create donation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Donation created successfully", "donation": donation})
}

// ListDonations handles GET requests to retrieve all donations
func ListDonations(c *gin.Context) {
	var donations []models.Donation

	if err := db.DB.Preload("Donor").Find(&donations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve donations"})
		return
	}

	c.JSON(http.StatusOK, donations)
}
