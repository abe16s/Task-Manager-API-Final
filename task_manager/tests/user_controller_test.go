package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/delivery/controllers"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserControllerSuite struct {
	suite.Suite
	controller   *controllers.UserController
	mockService  *mocks.UserServiceInterface
}

func (suite *UserControllerSuite) SetupTest() {
	suite.mockService = new(mocks.UserServiceInterface)
	suite.controller = &controllers.UserController{Service: suite.mockService}
}

// Test RegisterUser

func (suite *UserControllerSuite) TestRegisterUser_Success() {
	user := domain.User{Username: "testuser", Password: "password123"}
	suite.mockService.On("RegisterUser", &user).Return(&user, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	userJSON, _ := json.Marshal(user)
	c.Request, _ = http.NewRequest("POST", "/register", bytes.NewBuffer(userJSON))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.controller.RegisterUser(c)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "testuser")
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestRegisterUser_UsernameAlreadyExists() {
	user := domain.User{Username: "existinguser", Password: "password123"}
	suite.mockService.On("RegisterUser", &user).Return(nil, fmt.Errorf("username already exists"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	userJSON, _ := json.Marshal(user)
	c.Request, _ = http.NewRequest("POST", "/register", bytes.NewBuffer(userJSON))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.controller.RegisterUser(c)

	assert.Equal(suite.T(), http.StatusConflict, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "username already exists")
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestRegisterUser_InvalidJSON() {
	invalidJSON := "{invalid json"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("POST", "/register", bytes.NewBufferString(invalidJSON))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.controller.RegisterUser(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid JSON")
}

func (suite *UserControllerSuite) TestRegisterUser_ValidationErrors() {
	user := domain.User{Password: "password123"} // missing username
	suite.mockService.On("RegisterUser", &user).Return(nil, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	userJSON, _ := json.Marshal(user)
	c.Request, _ = http.NewRequest("POST", "/register", bytes.NewBuffer(userJSON))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.controller.RegisterUser(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "username is required.")
}

// Test Login

func (suite *UserControllerSuite) TestLogin_Success() {
	user := domain.User{Username: "testuser", Password: "password123"}
	token := "some-valid-token"
	suite.mockService.On("LoginUser", user).Return(token, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	userJSON, _ := json.Marshal(user)
	c.Request, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(userJSON))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.controller.Login(c)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "User logged in successfully")
	assert.Contains(suite.T(), w.Body.String(), token)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestLogin_UserNotFound() {
	user := domain.User{Username: "nonexistent", Password: "password123"}
	suite.mockService.On("LoginUser", user).Return("", fmt.Errorf("user not found"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	userJSON, _ := json.Marshal(user)
	c.Request, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(userJSON))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.controller.Login(c)

	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "user not found")
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestLogin_InvalidCredentials() {
	user := domain.User{Username: "testuser", Password: "wrongpassword"}
	suite.mockService.On("LoginUser", user).Return("", fmt.Errorf("invalid credentials"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	userJSON, _ := json.Marshal(user)
	c.Request, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(userJSON))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.controller.Login(c)

	assert.Equal(suite.T(), http.StatusUnauthorized, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "invalid credentials")
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestLogin_InvalidJSON() {
	invalidJSON := "{invalid json"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("POST", "/login", bytes.NewBufferString(invalidJSON))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.controller.Login(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "invalid character")
}

// Test PromoteUser

func (suite *UserControllerSuite) TestPromoteUser_Success() {
	username := "testuser"
	suite.mockService.On("PromoteUser", username).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("PUT", "/promote?username="+username, nil)

	suite.controller.PromoteUser(c)
	c.Writer.WriteHeaderNow()
	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestPromoteUser_InternalServerError() {
	username := "testuser"
	suite.mockService.On("PromoteUser", username).Return(fmt.Errorf("internal error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("PUT", "/promote?username="+username, nil)

	suite.controller.PromoteUser(c)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "internal error")
	suite.mockService.AssertExpectations(suite.T())
}

func TestUserControllerSuite(t *testing.T) {
	suite.Run(t, new(UserControllerSuite))
}
