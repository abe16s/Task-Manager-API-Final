package usecases

import "github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"

type UserRepoInterface interface {
	RegisterUser(user *domain.User) (*domain.User, error)
	GetUser(username string) (*domain.User, error)
	PromoteUser(username string) error
	Count() (int64, error)
}