package controllers

import (
	"net/http"

	"github.com/elishambadi/sharebite/db"
	"github.com/elishambadi/sharebite/models"
	"github.com/elishambadi/sharebite/services"
	"github.com/elishambadi/sharebite/utils"

	"github.com/gin-gonic/gin"
)

// CreateDonation handles POST requests to log a new food donation
func CreateDonation(c *gin.Context) {
	user, err := services.GetUserFromRequest(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting user from request",
		})
		c.Abort()
		return
	}

	var donation models.Donation
	if err := c.ShouldBindJSON(&donation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	donation.DonorID = user.ID

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

	c.JSON(http.StatusOK, gin.H{
		"message":   "Donations retrieved successfully",
		"donations": donations,
	})
}

func UploadDonationImage(c *gin.Context) {
	uploadDir := "./uploads/donations" // Directory to save uploaded files
	imageURL, err := utils.UploadFile(c, uploadDir)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_url": imageURL})
}

// Handling donation requests
