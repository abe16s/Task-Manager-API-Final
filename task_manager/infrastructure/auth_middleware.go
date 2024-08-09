package infrastructure

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(adminCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token, err := ValidateAuthHeader(authHeader)
		
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		// if adminCheck is true check for is_admin in the token by decoding
		if adminCheck {
			if !ValidateAdmin(token) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}