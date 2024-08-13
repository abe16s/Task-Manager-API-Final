package usecases

import "github.com/dgrijalva/jwt-go"

type JwtServiceInterface interface {
	GenerateToken(username string, isAdmin bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	ValidateAdmin(token *jwt.Token) bool
}