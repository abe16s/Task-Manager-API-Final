package infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
}

func (p *PasswordService) HashPassword(password string) (string, error) {
	hashedPassword, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if e != nil {
		return "", e
	}
	return string(hashedPassword), nil
}

func (p *PasswordService) ComparePassword(existingPassword string, userPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existingPassword), []byte(userPassword)) == nil
}
