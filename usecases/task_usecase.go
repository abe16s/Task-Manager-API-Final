package usecases

import (
	"errors"
	"strings"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
	"github.com/google/uuid"
)

type TaskServiceInterface interface {
	GetTasks() ([]domain.Task, error)
	GetTaskById(id uuid.UUID) (*domain.Task, error)
	UpdateTaskByID(id uuid.UUID, updatedTask domain.Task) error
	DeleteTask(id uuid.UUID) error
	AddTask(task domain.Task) (*domain.Task, error)
}

type TaskService struct {
	TaskRepo TaskRepoInterface
}


func (s *TaskService) GetTasks() ([]domain.Task, error) {
	tasks, err := s.TaskRepo.GetTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) GetTaskById(id uuid.UUID) (*domain.Task, error) {
	task, err := s.TaskRepo.GetTaskById(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) UpdateTaskByID(id uuid.UUID, updatedTask domain.Task) error {
	if strings.ToLower(updatedTask.Status) != "in progress" && strings.ToLower(updatedTask.Status) != "completed" && strings.ToLower(updatedTask.Status) != "pending" {
		return errors.New("status error")
	}
	
	err := s.TaskRepo.UpdateTaskByID(id, updatedTask)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) DeleteTask(id uuid.UUID) error {
	err := s.TaskRepo.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) AddTask(task domain.Task) (*domain.Task, error) {
	if strings.ToLower(task.Status) != "in progress" && strings.ToLower(task.Status) != "completed" && strings.ToLower(task.Status) != "pending" {
		return nil, errors.New("status error")
	}
	task.ID = uuid.New()
	newTask, err := s.TaskRepo.AddTask(task)
	if err != nil {
		return nil, err
	}
	return newTask, nil
}