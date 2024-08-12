package infrastructure

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func AuthMiddleware(adminCheck bool) gin.HandlerFunc {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// get jwt secret from env
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	var jwtservice JwtServiceInterface = &JwtService{JwtSecret: jwtSecret}

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwtservice.ValidateToken(authParts[1])

		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// if adminCheck is true check for is_admin in the token by decoding
		if adminCheck {
			if !jwtservice.ValidateAdmin(token) {
				c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
