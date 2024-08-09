package usecases

import (
	"errors"
	"strings"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
	"github.com/google/uuid"
)

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

func (s *TaskService) UpdateTaskByID(id uuid.UUID, updatedTask domain.Task) (*domain.Task, error) {
	task, err := s.TaskRepo.UpdateTaskByID(id, updatedTask)
	if err != nil {
		return nil, err
	}
	return task, nil
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
	newTask, err := s.TaskRepo.AddTask(task)
	if err != nil {
		return nil, err
	}
	return newTask, nil
}