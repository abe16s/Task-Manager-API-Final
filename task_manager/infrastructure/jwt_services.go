package infrastructure

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var jwtSecret []byte

func init() {
	//load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// get jwt secret from env
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}


// validate authHeader
func ValidateAuthHeader(authHeader string) (*jwt.Token, error) {
	if authHeader == "" {
		return nil, errors.New("authorization header is required")
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return nil, errors.New("invalid authorization header")
	}

	token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})	

	if err != nil || !token.Valid {
		return nil, errors.New("invalid JWT")
	}
	
	return token, err
}


func ValidateAdmin(token *jwt.Token) bool {
	claims, ok := token.Claims.(jwt.MapClaims)
	return ok && claims["is_admin"].(bool)
}