package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/elishambadi/sharebite/db"
	"github.com/elishambadi/sharebite/models"
	"github.com/elishambadi/sharebite/utils"
	"github.com/gin-gonic/gin"
)

func CheckUserRole(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")

	if tokenStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No authorization headers."})
		c.Abort()
		return
	}
	splitString := strings.Split(tokenStr, " ")
	if len(splitString) == 2 {
		tokenStr = splitString[1]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Malformed Auth Headers"})
		c.Abort()
		return
	}
	log.Println("Token in CheckUserRole: ", tokenStr)

	token, err := utils.ValidateJWT(tokenStr)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid token: %s", err)})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}

	roles := claims["roles"].([]interface{})
	// Check if the user has the required role
	hasAdminRole := false
	for _, role := range roles {
		if role == "admin" {
			hasAdminRole = true
			break
		}
	}

	if !hasAdminRole {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
		return
	}

	var user models.User
	if err := db.DB.Where("api_token = ?", tokenStr).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}

	// Store user in context if needed
	c.Set("user", user)
	c.Next()
}
