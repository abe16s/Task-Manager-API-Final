package usecases

import (
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
)

type UserService struct {
	UserRepo UserRepoInterface
}

// register new user with unique username and password
func (s *UserService) RegisterUser(user domain.User) (*domain.User, error) {
	count, err := s.UserRepo.Count()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		user.IsAdmin = true
	}
	u, err := s.UserRepo.RegisterUser(user)

	if err != nil {
		return nil, err
	}

	return u, nil
}


// login user 
func (s *UserService) LoginUser(user domain.User) (string, error) {
	token, err := s.UserRepo.LoginUser(user)
	if err != nil {
		return "", err
	}

	return token, nil
}


// promote user to admin
func (s *UserService) PromoteUser(username string) error {
	return s.UserRepo.PromoteUser(username)
}