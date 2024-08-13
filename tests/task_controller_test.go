package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/delivery/controllers"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TaskControllerSuite struct {
	suite.Suite
	controller *controllers.TaskController
	mockService *mocks.TaskServiceInterface
}

func (suite *TaskControllerSuite) SetupTest() {
	suite.mockService = new(mocks.TaskServiceInterface)
	suite.controller = &controllers.TaskController{
		Service: suite.mockService,
	}
}

func (suite *TaskControllerSuite) TestGetTasks_Success() {
	tasks := []domain.Task{
		{ID: uuid.New(), Title: "Task 1", Description: "Description 1", Status: "pending", DueDate: time.Now().Truncate(0)},
		{ID: uuid.New(), Title: "Task 2", Description: "Description 2", Status: "done", DueDate: time.Now().Truncate(0)},
	}

	suite.mockService.On("GetTasks").Return(tasks, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	suite.controller.GetTasks(c)

	suite.Equal(http.StatusOK, w.Code)
	var gotTasks []domain.Task
	err := json.Unmarshal(w.Body.Bytes(), &gotTasks)
	suite.NoError(err)
	suite.Equal(tasks, gotTasks)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestGetTaskById_Success() {
	id := uuid.New()
	task := &domain.Task{ID: id, Title: "Task 1", Description: "Description 1", Status: "pending", DueDate: time.Now().Truncate(0)}

	suite.mockService.On("GetTaskById", id).Return(task, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: id.String()}}

	suite.controller.GetTaskById(c)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	var gotTask domain.Task
	err := json.Unmarshal(w.Body.Bytes(), &gotTask)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), *task, gotTask)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestGetTaskById_NotFound() {
	id := uuid.New()

	suite.mockService.On("GetTaskById", id).Return(nil, errors.New("Task Not Found"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: id.String()}}

	suite.controller.GetTaskById(c)

	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestGetTaskById_InvalidUUID() {
    invalidUUID := "invalid-uuid"

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Params = gin.Params{gin.Param{Key: "id", Value: invalidUUID}}

    suite.controller.GetTaskById(c)

    assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
    assert.Contains(suite.T(), w.Body.String(), "invalid task ID")
}


func (suite *TaskControllerSuite) TestAddTask_Success() {
	task := &domain.Task{ID: uuid.New(), Title: "New Task", Description: "New Description", Status: "pending", DueDate: time.Now()}

	suite.mockService.On("AddTask", mock.AnythingOfType("domain.Task")).Return(task, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("POST", "/tasks", bytes.NewBufferString(`{"title": "New Task", "description": "New Description", "status": "pending", "due_date": "`+task.DueDate.Format(time.RFC3339)+`"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.controller.AddTask(c)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestAddTask_InvalidJSON() {
    invalidJSON := "{invalid json"

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    c.Request, _ = http.NewRequest("POST", "/tasks", bytes.NewBufferString(invalidJSON))
    c.Request.Header.Set("Content-Type", "application/json")

    suite.controller.AddTask(c)

    assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	fmt.Println(w.Body.String())
    assert.Contains(suite.T(), w.Body.String(), "Invalid JSON")  
}

func (suite *TaskControllerSuite) TestAddTask_ValidationErrors() {
    invalidTask := domain.Task{Status: "completed", DueDate: time.Now()}
    taskJSON, _ := json.Marshal(invalidTask)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    c.Request, _ = http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
    c.Request.Header.Set("Content-Type", "application/json")

    suite.controller.AddTask(c)

    assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
    assert.Contains(suite.T(), w.Body.String(), "Title is required.")
    assert.Contains(suite.T(), w.Body.String(), "Description is required.")
}


func (suite *TaskControllerSuite) TestUpdateTaskByID_Success() {
	id := uuid.New()
	task := domain.Task{ID: id, Title: "Updated Task", Description: "Updated Description", Status: "completed", DueDate: time.Now()}

	suite.mockService.On("UpdateTaskByID", id, mock.AnythingOfType("domain.Task")).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: id.String()}}

	taskJSON, _ := json.Marshal(task)
	c.Request, _ = http.NewRequest("PUT", "/tasks/"+id.String(), bytes.NewBuffer(taskJSON))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.controller.UpdateTaskByID(c)
	c.Writer.WriteHeaderNow()
	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestUpdateTaskByID_InvalidUUID() {
    invalidUUID := "invalid-uuid"

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Params = gin.Params{gin.Param{Key: "id", Value: invalidUUID}}

    c.Request, _ = http.NewRequest("PUT", "/tasks/"+invalidUUID, nil)

    suite.controller.UpdateTaskByID(c)

    assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
    assert.Contains(suite.T(), w.Body.String(), "invalid task ID")
}

func (suite *TaskControllerSuite) TestUpdateTaskByID_InvalidJSON() {
    id := uuid.New()
    invalidJSON := "{invalid json"

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Params = gin.Params{gin.Param{Key: "id", Value: id.String()}}

    c.Request, _ = http.NewRequest("PUT", "/tasks/"+id.String(), bytes.NewBufferString(invalidJSON))
    c.Request.Header.Set("Content-Type", "application/json")

    suite.controller.UpdateTaskByID(c)

    assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
    assert.Contains(suite.T(), w.Body.String(), "Invalid JSON") 
}

func (suite *TaskControllerSuite) TestUpdateTaskByID_ValidationErrors() {
    id := uuid.New()
    // Missing required fields (e.g., Title, Description)
    invalidTask := domain.Task{Status: "completed", DueDate: time.Now()}
    taskJSON, _ := json.Marshal(invalidTask)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Params = gin.Params{gin.Param{Key: "id", Value: id.String()}}

    c.Request, _ = http.NewRequest("PUT", "/tasks/"+id.String(), bytes.NewBuffer(taskJSON))
    c.Request.Header.Set("Content-Type", "application/json")

    suite.controller.UpdateTaskByID(c)

    assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
    assert.Contains(suite.T(), w.Body.String(), "Title is required.")
    assert.Contains(suite.T(), w.Body.String(), "Description is required.")
}


func (suite *TaskControllerSuite) TestDeleteTask_Success() {
	id := uuid.New()

	suite.mockService.On("DeleteTask", id).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: id.String()}}

	suite.controller.DeleteTask(c)
	c.Writer.WriteHeaderNow()

	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestDeleteTask_InvalidUUID() {
    invalidUUID := "invalid-uuid"

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Params = gin.Params{gin.Param{Key: "id", Value: invalidUUID}}

    suite.controller.DeleteTask(c)

    assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
    assert.Contains(suite.T(), w.Body.String(), "invalid task ID")
}


func TestTaskControllerSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerSuite))
}
