package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/models"
	// "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"    
)

var tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}


type TaskService struct {
	Client *mongo.Client 
}



func (s *TaskService) GetTasks() []models.Task {
	return tasks
}

func (s *TaskService) GetTaskById(id string) (*models.Task, error) {
	for _, val := range tasks {
		if val.ID == id {
			return &val, nil
		}
	}
	return nil, errors.New("task not found")
}

func (s *TaskService) UpdateTaskByID(id string, updateTask models.Task) (*models.Task, error) {
	task, err := s.GetTaskById(id)
	if err != nil {
		return nil, err
	}

	task.Title = updateTask.Title

	if updateTask.Status == "Pending" || updateTask.Status == "In Progress" || updateTask.Status == "Completed" {
		task.Status = updateTask.Status
	} else if updateTask.Status != "" {
		return nil, errors.New("status must be 'Pending' or 'In Progress' or 'Completed'")
	}

	task.Description = updateTask.Description
	task.DueDate = updateTask.DueDate

	return task, nil
}

func (s *TaskService) DeleteTask(id string) (*models.Task, error) {
	for i, val := range tasks {
		if val.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return &val, nil
		}
	}
	return nil, errors.New("task not found")
}

func (s *TaskService) AddTask(task models.Task) *models.Task {
	newId := strconv.Itoa(len(tasks) + 1)
	task.ID = newId
	if task.Status == "" {
		task.Status = "Pending"
	}
	tasks = append(tasks, task)
	return &tasks[len(tasks)-1]
}
