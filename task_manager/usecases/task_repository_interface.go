package usecases

import (
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
	"github.com/google/uuid"
)

type TaskRepoInterface interface {
	GetTasks() ([]domain.Task, error)
	GetTaskById(id uuid.UUID) (*domain.Task, error)
	UpdateTaskByID(id uuid.UUID, updatedTask domain.Task) (*domain.Task, error)
	DeleteTask(id uuid.UUID) error
	AddTask(task domain.Task) (*domain.Task, error)
}