package usecases

import (
	"errors"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/infrastructure"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepo UserRepoInterface
	PasswordService infrastructure.PasswordServiceInterface
	JwtService infrastructure.JwtServiceInterface
}

// register new user with unique username and password
func (s *UserService) RegisterUser(user *domain.User) (*domain.User, error) {
	count, err := s.UserRepo.Count()
	if err != nil {
		return nil, err
	}

	user.ID = uuid.New()

	if count == 0 {
		user.IsAdmin = true
	}

	hashedPassword, err := s.PasswordService.HashPassword(user.Password)
    if err != nil {
		return nil, err
    }

    user.Password = hashedPassword

	u, err := s.UserRepo.RegisterUser(user)

	if err != nil {
		return nil, err
	}

	return u, nil
}


// login user 
func (s *UserService) LoginUser(user domain.User) (string, error) {
	existingUser, err := s.UserRepo.GetUser(user.Username)
	if err != nil {
		return "", err
	}

	match := s.PasswordService.ComparePassword(existingUser.Password, user.Password)
	if !match {
		return "", errors.New("invalid credentials")
	}

	// generate token
	jwtToken, err := s.JwtService.GenerateToken(existingUser.Username, existingUser.IsAdmin)
	if err != nil {
		return "", errors.New("internal server error")
	}

	return jwtToken, nil
}


// promote user to admin
func (s *UserService) PromoteUser(username string) error {
	return s.UserRepo.PromoteUser(username)
}