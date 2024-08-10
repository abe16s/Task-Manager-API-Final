package infrastructure

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
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

func GenerateToken(username string, isAdmin bool) (string, error) {
	expirationTime := time.Now().Add(20 * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"is_admin":  isAdmin,
		"exp":      expirationTime,
	})
	jwtToken, e := token.SignedString(jwtSecret)

	if e != nil {
		return "", errors.New("can't sign token")
	}

	return jwtToken, nil
}

// validate token
func ValidateToken(token string) (*jwt.Token, error) {
	jwtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})	

	if err != nil || !jwtoken.Valid {
		return nil, errors.New("invalid JWT")
	}
	
	return jwtoken, err
}


func ValidateAdmin(token *jwt.Token) bool {
	claims, ok := token.Claims.(jwt.MapClaims)
	return ok && claims["is_admin"].(bool)
}