package middleware

import (
	"depo-bangunan/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, _ := c.Get("currentUser")
		userLoggedIn, _ := currentUser.(*models.UserLoggedIn)

		if userLoggedIn.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "Only admin can do"})
			return
		}
		
		c.Next()
	}
}
