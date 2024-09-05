package middleware

import (
	"depo-bangunan/config"
	"depo-bangunan/helpers"
	"depo-bangunan/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		access_token := c.Request.Header.Get("Authorization")
		if access_token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "You are not logged in"})
			return
		}

		sub, err := helpers.ValidateToken(access_token, config.ENV.AccessTokenPublicKey)
		subMap, _ := sub.(map[string]interface{})
		email, _ := subMap["email"].(string)
		role, _ := subMap["role"].(string)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": err.Error()})
			return
		}

		var result *gorm.DB
		var userId int
		var user models.User
		
		result = config.DB.First(&user, "email = ?", email)
		userId = int(user.ID)

		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "The user belonging to this token no logger exists"})
			return
		}

		userLoggedId := &models.UserLoggedIn{
			ID: userId,
			Email: email,
			Role: role,
		}

		c.Set("currentUser", userLoggedId)
		c.Next()
	}
}
