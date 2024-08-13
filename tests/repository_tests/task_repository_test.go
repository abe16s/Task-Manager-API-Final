package tests

import (
	"context"
	"testing"
	"time"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepositorySuite struct {
	suite.Suite
	client     *mongo.Client
	repo       *repositories.TaskRepository
	collection *mongo.Collection
}

func (suite *TaskRepositorySuite) SetupSuite() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.client = client
	suite.collection = client.Database("test_db").Collection("tasks")
	suite.repo = repositories.NewTaskRepository(client, "test_db", "tasks")
}

func (suite *TaskRepositorySuite) TearDownSuite() {
	err := suite.client.Disconnect(context.Background())
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *TaskRepositorySuite) SetupTest() {
	// Clear the collection before each test
	err := suite.collection.Drop(context.Background())
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *TaskRepositorySuite) TearDownTest() {
	// Additional cleanup if needed
	_, err := suite.collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *TaskRepositorySuite) TestAddTask() {
	task := domain.Task{
		ID:          uuid.New(),
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
		DueDate:     time.Now().UTC(),
	}

	addedTask, err := suite.repo.AddTask(task)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), task, *addedTask)
}

func (suite *TaskRepositorySuite) TestGetTasks() {
	task1 := domain.Task{
		ID:          uuid.New(),
		Title:       "Test Task 1",
		Description: "Test Description 1",
		Status:      "pending",
		DueDate:     time.Now().UTC(),
	}
	task2 := domain.Task{
		ID:          uuid.New(),
		Title:       "Test Task 2",
		Description: "Test Description 2",
		Status:      "completed",
		DueDate:     time.Now().UTC(),
	}

	_, err := suite.repo.AddTask(task1)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.AddTask(task2)
	assert.NoError(suite.T(), err)

	tasks, err := suite.repo.GetTasks()
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), tasks, 2)
}

func (suite *TaskRepositorySuite) TestGetTaskById() {
	task := domain.Task{
		ID:          uuid.New(),
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
		DueDate:     time.Now().UTC(),
	}

	_, err := suite.repo.AddTask(task)
	assert.NoError(suite.T(), err)

	foundTask, err := suite.repo.GetTaskById(task.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), task.ID, foundTask.ID)
}

func (suite *TaskRepositorySuite) TestUpdateTaskByID() {
	task := domain.Task{
		ID:          uuid.New(),
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
		DueDate:     time.Now().UTC(),
	}

	_, err := suite.repo.AddTask(task)
	assert.NoError(suite.T(), err)

	updatedTask := domain.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "completed",
		DueDate:     time.Now().UTC(),
	}

	err = suite.repo.UpdateTaskByID(task.ID, updatedTask)
	assert.NoError(suite.T(), err)

	foundTask, err := suite.repo.GetTaskById(task.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), updatedTask.Title, foundTask.Title)
	assert.Equal(suite.T(), updatedTask.Description, foundTask.Description)
	assert.Equal(suite.T(), updatedTask.Status, foundTask.Status)
}

func (suite *TaskRepositorySuite) TestDeleteTask() {
	task := domain.Task{
		ID:          uuid.New(),
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
		DueDate:     time.Now().UTC(),
	}

	_, err := suite.repo.AddTask(task)
	assert.NoError(suite.T(), err)

	err = suite.repo.DeleteTask(task.ID)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetTaskById(task.ID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "task Not Found", err.Error())
}

func TestTaskRepositorySuite(t *testing.T) {
	suite.Run(t, new(TaskRepositorySuite))
}