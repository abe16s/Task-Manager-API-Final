package repository_tests

import (
	"context"
	"testing"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositorySuite struct {
	suite.Suite
	client         *mongo.Client
	collection *mongo.Collection
	repo       *repositories.UserRepository
}

func (suite *UserRepositorySuite) SetupSuite() {
	// Connect to the MongoDB instance
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		suite.T().Fatal(err)
	}

	// Set up a test database
	suite.client = client
	suite.collection = client.Database("test_db").Collection("users")
	suite.repo = repositories.NewUserRepository(client, "test_db", "users")
}

func (suite *UserRepositorySuite) TearDownSuite() {
	// Disconnect the client
	err := suite.client.Disconnect(context.Background())
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *UserRepositorySuite) TearDownTest() {
	// Clear the collection after each test
	_, err := suite.collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *UserRepositorySuite) TestRegisterUser_Success() {
	user := &domain.User{
		ID:       uuid.New(),
		Username: "testuser",
		Password: "testpassword",
		IsAdmin:  false,
	}

	createdUser, err := suite.repo.RegisterUser(user)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Username, createdUser.Username)
	assert.Equal(suite.T(), user.Password, createdUser.Password)
}

func (suite *UserRepositorySuite) TestRegisterUser_UsernameExists() {
	user := &domain.User{
		ID:       uuid.New(),
		Username: "testuser",
		Password: "testpassword",
		IsAdmin:  false,
	}

	_, err := suite.repo.RegisterUser(user)
	assert.NoError(suite.T(), err)

	user.ID = uuid.New()
	// Try to register the same user again
	_, err = suite.repo.RegisterUser(user)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "username already exists", err.Error())
}

func (suite *UserRepositorySuite) TestGetUser_Success() {
	user := &domain.User{
		ID:       uuid.New(),
		Username: "testuser",
		Password: "testpassword",
		IsAdmin:  false,
	}

	_, err := suite.repo.RegisterUser(user)
	assert.NoError(suite.T(), err)

	// Retrieve the user
	retrievedUser, err := suite.repo.GetUser(user.Username)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Username, retrievedUser.Username)
	assert.Equal(suite.T(), user.Password, retrievedUser.Password)
}

func (suite *UserRepositorySuite) TestGetUser_NotFound() {
	_, err := suite.repo.GetUser("nonexistentuser")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "user not found", err.Error())
}

func (suite *UserRepositorySuite) TestPromoteUser_Success() {
	user := &domain.User{
		ID:       uuid.New(),
		Username: "testuser",
		Password: "testpassword",
		IsAdmin:  false,
	}

	_, err := suite.repo.RegisterUser(user)
	assert.NoError(suite.T(), err)

	// Promote the user
	err = suite.repo.PromoteUser(user.Username)
	assert.NoError(suite.T(), err)

	// Retrieve the user and check the isAdmin field
	updatedUser, err := suite.repo.GetUser(user.Username)
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), updatedUser.IsAdmin)
}

func (suite *UserRepositorySuite) TestPromoteUser_NotFound() {
	err := suite.repo.PromoteUser("nonexistentuser")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "username not found", err.Error())
}

// Test Suite Execution
func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
