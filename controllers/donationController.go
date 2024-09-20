package controllers

import (
	"net/http"

	"github.com/elishambadi/sharebite/models"
	"github.com/elishambadi/sharebite/services"

	"github.com/gin-gonic/gin"
)

var donationService *services.DonationService

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

	if err := donationService.CreateDonation(&donation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create donation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Donation created successfully", "donation": donation})
}

// ListDonations handles GET requests to retrieve all donations
func ListDonations(c *gin.Context) {
	donations, err := donationService.ListDonations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve donations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Donations retrieved successfully",
		"donations": donations,
	})
}

// UploadDonationImage handles image uploads for donations
func UploadDonationImage(c *gin.Context) {
	uploadDir := "./uploads/donations" // Directory to save uploaded files
	imageURL, err := donationService.UploadDonationImage(uploadDir, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_url": imageURL})
}

// CreateDonationRequest handles the creation of a donation request
func CreateDonationRequest(c *gin.Context) {
	var request models.DonationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Get the user from the request to set DonorID
	user, err := services.GetUserFromRequest(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting user from request"})
		return
	}

	request.RecipientID = user.ID // Any user can be a donor

	if err := donationService.CreateDonationRequest(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Request created successfully", "request": request})
}

// UpdateDonationRequestStatus handles updating the status of a donation request
func UpdateDonationRequestStatus(c *gin.Context) {
	requestID := c.Param("id")
	var input struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	// Define valid statuses with descriptions
	validStatuses := map[string]string{
		"PENDING":   "The request is pending.",
		"APPROVED":  "The request has been approved.",
		"COMPLETED": "The donation has been completed.",
		"REJECTED":  "The request has been rejected.",
	}

	// Check if the input status is valid
	description, exists := validStatuses[input.Status]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status value"})
		return
	}

	// Get the user from the request
	user, err := services.GetUserFromRequest(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting user from request"})
		return
	}

	// Fetch the donation request
	request, err := donationService.GetDonationRequestByID(requestID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Donation request not found"})
		return
	}

	// Check if the user is the donor
	if user.ID == request.Donation.DonorID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You cannot update your own donation request status"})
		return
	}

	// Update the donation request status
	if err := donationService.UpdateDonationRequestStatus(requestID, input.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update request status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Request status updated successfully", "description": description})
}

// ListDonationRequests retrieves all donation requests
func ListDonationRequests(c *gin.Context) {
	requests, err := donationService.ListDonationRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve requests"})
		return
	}

	c.JSON(http.StatusOK, requests)
}
