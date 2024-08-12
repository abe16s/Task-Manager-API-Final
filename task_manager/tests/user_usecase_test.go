package tests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/tests/mocks"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/usecases"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	service        *usecases.UserService
	mockJwtService *mocks.JwtServiceInterface
	mockPwdService *mocks.PasswordServiceInterface
	mockUserRepo   *mocks.UserRepoInterface
}

// Setup test environment
func (suite *UserServiceTestSuite) SetupTest() {
	suite.mockJwtService = new(mocks.JwtServiceInterface)
	suite.mockPwdService = new(mocks.PasswordServiceInterface)
	suite.mockUserRepo = new(mocks.UserRepoInterface)
	suite.service = &usecases.UserService{
		UserRepo:        suite.mockUserRepo,
		PasswordService: suite.mockPwdService,
		JwtService:      suite.mockJwtService,
	}
}

// Test RegisterUser with a new user
func (suite *UserServiceTestSuite) TestRegisterUser_NewUser() {
	user := domain.User{
		Username: "testuser",
		Password: "password123",
	}

	// Mocking the Count method to return 0 (first user)
	suite.mockUserRepo.On("Count").Return(int64(0), nil)
	// Mocking the HashPassword method to return a hashed password
	suite.mockPwdService.On("HashPassword", "password123").Return("hashedpassword123", nil)
	// Mocking the RegisterUser method to return the registered user
	suite.mockUserRepo.On("RegisterUser", mock.AnythingOfType("*domain.User")).Return(&user, nil)

	registeredUser, err := suite.service.RegisterUser(&user)
	fmt.Println(registeredUser)

	suite.NoError(err)
	suite.NotNil(registeredUser)
	suite.Equal("testuser", registeredUser.Username)
	suite.Equal("hashedpassword123", registeredUser.Password)
	suite.True(registeredUser.IsAdmin)

	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockPwdService.AssertExpectations(suite.T())
}

// Test RegisterUser with an existing user
func (suite *UserServiceTestSuite) TestRegisterUser_ExistingUser() {
	user := domain.User{
		Username: "testuser",
		Password: "password123",
	}

	// Mocking the Count method to return 1 (not the first user)
	suite.mockUserRepo.On("Count").Return(int64(1), nil)
	// Mocking the HashPassword method to return a hashed password
	suite.mockPwdService.On("HashPassword", "password123").Return("hashedpassword123", nil)
	// Mocking the RegisterUser method to return the registered user
	suite.mockUserRepo.On("RegisterUser", mock.AnythingOfType("*domain.User")).Return(&user, nil)

	registeredUser, err := suite.service.RegisterUser(&user)

	suite.NoError(err)
	suite.NotNil(registeredUser)
	suite.Equal("testuser", registeredUser.Username)
	suite.Equal("hashedpassword123", registeredUser.Password)
	suite.False(registeredUser.IsAdmin) // Not the first user, hence not an admin

	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockPwdService.AssertExpectations(suite.T())
}

// Test LoginUser with valid credentials
func (suite *UserServiceTestSuite) TestLoginUser_ValidCredentials() {
	user := domain.User{
		Username: "testuser",
		Password: "password123",
	}

	// Mocking the GetUser method to return the existing user
	suite.mockUserRepo.On("GetUser", "testuser").Return(&user, nil)
	// Mocking the ComparePassword method to return true
	suite.mockPwdService.On("ComparePassword", user.Password, "password123").Return(true)
	// Mocking the GenerateToken method to return a JWT token
	suite.mockJwtService.On("GenerateToken", "testuser", false).Return("valid.jwt.token", nil)

	token, err := suite.service.LoginUser(user)

	suite.NoError(err)
	suite.Equal("valid.jwt.token", token)

	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockPwdService.AssertExpectations(suite.T())
	suite.mockJwtService.AssertExpectations(suite.T())
}

// Test LoginUser with invalid credentials
func (suite *UserServiceTestSuite) TestLoginUser_InvalidCredentials() {
	user := domain.User{
		Username: "testuser",
		Password: "wrongpassword",
	}

	// Mocking the GetUser method to return the existing user
	suite.mockUserRepo.On("GetUser", "testuser").Return(&user, nil)
	// Mocking the ComparePassword method to return false
	suite.mockPwdService.On("ComparePassword", user.Password, "wrongpassword").Return(false)

	token, err := suite.service.LoginUser(user)

	suite.EqualError(err, "invalid credentials")
	suite.Empty(token)

	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockPwdService.AssertExpectations(suite.T())
}

// Test PromoteUser
func (suite *UserServiceTestSuite) TestPromoteUser() {
	// Mocking the PromoteUser method to return nil
	suite.mockUserRepo.On("PromoteUser", "testuser").Return(nil)

	err := suite.service.PromoteUser("testuser")

	suite.NoError(err)

	suite.mockUserRepo.AssertExpectations(suite.T())
}

// Test PromoteUser with an error
func (suite *UserServiceTestSuite) TestPromoteUser_Error() {
	// Mocking the PromoteUser method to return an error
	suite.mockUserRepo.On("PromoteUser", "testuser").Return(errors.New("promotion failed"))

	err := suite.service.PromoteUser("testuser")

	suite.EqualError(err, "promotion failed")

	suite.mockUserRepo.AssertExpectations(suite.T())
}

// Run the test suite
func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}