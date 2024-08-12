package infrastructure

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtServiceInterface interface {
	GenerateToken(username string, isAdmin bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	ValidateAdmin(token *jwt.Token) bool
}

type JwtService struct {
	JwtSecret []byte
}

func (j *JwtService) GenerateToken(username string, isAdmin bool) (string, error) {
	expirationTime := time.Now().Add(20 * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"is_admin": isAdmin,
		"exp":      expirationTime,
	})
	jwtToken, e := token.SignedString(j.JwtSecret)

	if e != nil {
		return "", errors.New("can't sign token")
	}

	return jwtToken, nil
}

// validate token
func (j *JwtService) ValidateToken(token string) (*jwt.Token, error) {
	jwtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return j.JwtSecret, nil
	})

	if err != nil || !jwtoken.Valid {
		return nil, errors.New("invalid JWT")
	}

	return jwtoken, err
}

func (j *JwtService) ValidateAdmin(token *jwt.Token) bool {
	claims, ok := token.Claims.(jwt.MapClaims)
	return ok && claims["is_admin"].(bool)
}
