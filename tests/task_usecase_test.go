package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/tests/mocks"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// TaskServiceTestSuite defines the test suite for TaskService
type TaskServiceTestSuite struct {
	suite.Suite
	service  *usecases.TaskService
	mockRepo *mocks.TaskRepoInterface
}

// SetupTest sets up the test environment before each test
func (suite *TaskServiceTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.TaskRepoInterface)
	suite.service = &usecases.TaskService{TaskRepo: suite.mockRepo}
}

// TestGetTasks tests the GetTasks method
func (suite *TaskServiceTestSuite) TestGetTasks() {
	mockTasks := []domain.Task{
		{ID: uuid.New(), Title: "Test Task 1", Status: "pending", Description: "Test Description 1", DueDate: time.Now().UTC()},
		{ID: uuid.New(), Title: "Test Task 2", Status: "completed",  Description: "Test Description 2", DueDate: time.Now().UTC().Add(24 * time.Hour)},
	}

	suite.mockRepo.On("GetTasks").Return(mockTasks, nil)

	tasks, err := suite.service.GetTasks()

	suite.NoError(err)
	suite.Equal(mockTasks, tasks)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestGetTaskById tests the GetTaskById method
func (suite *TaskServiceTestSuite) TestGetTaskById() {
	taskID := uuid.New()
	mockTask := &domain.Task{ID: taskID, Title: "Test Task", Status: "pending", Description: "Test Description", DueDate: time.Now().UTC()}

	suite.mockRepo.On("GetTaskById", taskID).Return(mockTask, nil)

	task, err := suite.service.GetTaskById(taskID)

	suite.NoError(err)
	suite.Equal(mockTask, task)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestGetTaskById_InvalidID tests the GetTaskById method with an invalid ID
func (suite *TaskServiceTestSuite) TestGetTaskById_InvalidID() {
	invalidID := uuid.New()

	suite.mockRepo.On("GetTaskById", invalidID).Return(nil, errors.New("task not found"))

	task, err := suite.service.GetTaskById(invalidID)

	suite.Nil(task)
	suite.EqualError(err, "task not found")
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestUpdateTaskByID_ValidStatus tests the UpdateTaskByID method with a valid status
func (suite *TaskServiceTestSuite) TestUpdateTaskByID_ValidStatus() {
	taskID := uuid.New()
	updatedTask := domain.Task{ID: taskID, Title: "Updated Task", Status: "in progress", Description: "Updated Description", DueDate: time.Now().UTC()}

	suite.mockRepo.On("UpdateTaskByID", taskID, updatedTask).Return(nil)

	err := suite.service.UpdateTaskByID(taskID, updatedTask)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestUpdateTaskByID_InvalidStatus tests the UpdateTaskByID method with an invalid status
func (suite *TaskServiceTestSuite) TestUpdateTaskByID_InvalidStatus() {
	taskID := uuid.New()
	updatedTask := domain.Task{ID: taskID, Title: "Updated Task", Status: "unknown",  Description: "Updated Description", DueDate: time.Now().UTC()}

	err := suite.service.UpdateTaskByID(taskID, updatedTask)

	suite.EqualError(err, "status error")
	suite.mockRepo.AssertNotCalled(suite.T(), "UpdateTaskByID", taskID, updatedTask)
}

// TestUpdateTaskByID_InvalidID tests the UpdateTaskByID method with an invalid ID
func (suite *TaskServiceTestSuite) TestUpdateTaskByID_InvalidID() {
	invalidID := uuid.New()
	updatedTask := domain.Task{ID: invalidID, Title: "Updated Task", Status: "in progress", Description: "Updated Description", DueDate: time.Now().UTC()}

	suite.mockRepo.On("UpdateTaskByID", invalidID, updatedTask).Return(errors.New("task not found"))

	err := suite.service.UpdateTaskByID(invalidID, updatedTask)

	suite.EqualError(err, "task not found")
	suite.mockRepo.AssertExpectations(suite.T())
}


// TestDeleteTask tests the DeleteTask method
func (suite *TaskServiceTestSuite) TestDeleteTask() {
	taskID := uuid.New()

	suite.mockRepo.On("DeleteTask", taskID).Return(nil)

	err := suite.service.DeleteTask(taskID)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestDeleteTask_InvalidID tests the DeleteTask method with an invalid ID
func (suite *TaskServiceTestSuite) TestDeleteTask_InvalidID() {
	invalidID := uuid.New()

	suite.mockRepo.On("DeleteTask", invalidID).Return(errors.New("task not found"))

	err := suite.service.DeleteTask(invalidID)

	suite.EqualError(err, "task not found")
	suite.mockRepo.AssertExpectations(suite.T())
}


// TestAddTask_ValidStatus tests the AddTask method with a valid status
func (suite *TaskServiceTestSuite) TestAddTask_ValidStatus() {
	task := domain.Task{Title: "New Task", Status: "pending", Description: "Updated Description", DueDate: time.Now().UTC()}

	suite.mockRepo.On("AddTask", mock.AnythingOfType("domain.Task")).Return(&task, nil)

	newTask, err := suite.service.AddTask(task)

	suite.NoError(err)
	suite.Equal(task.Title, newTask.Title)
	suite.Equal(task.Status, newTask.Status)
	suite.Equal(task.Description, newTask.Description)
	suite.Equal(task.DueDate, newTask.DueDate)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestAddTask_InvalidStatus tests the AddTask method with an invalid status
func (suite *TaskServiceTestSuite) TestAddTask_InvalidStatus() {
	task := domain.Task{Title: "New Task", Status: "unknown"}

	newTask, err := suite.service.AddTask(task)

	suite.Nil(newTask)
	suite.EqualError(err, "status error")
	suite.mockRepo.AssertNotCalled(suite.T(), "AddTask", task)
}

// TestSuite entry point
func TestTaskServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TaskServiceTestSuite))
}