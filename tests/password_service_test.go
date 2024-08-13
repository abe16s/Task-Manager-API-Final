package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/infrastructure"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/usecases"
)

type PasswordServiceSuite struct {
	suite.Suite
	service usecases.PasswordServiceInterface
}

func (suite *PasswordServiceSuite) SetupTest() {
	suite.service = &infrastructure.PasswordService{}
}

// Test HashPassword

func (suite *PasswordServiceSuite) TestHashPassword_Success() {
	password := "mySecureP@ssw0rd"
	hashedPassword, err := suite.service.HashPassword(password)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), hashedPassword)

	// Verify that the hashed password is different from the plain password
	assert.NotEqual(suite.T(), password, hashedPassword)
}

// Test ComparePassword

func (suite *PasswordServiceSuite) TestComparePassword_Success() {
	password := "mySecureP@ssw0rd"
	hashedPassword, _ := suite.service.HashPassword(password)

	isMatch := suite.service.ComparePassword(hashedPassword, password)
	assert.True(suite.T(), isMatch)
}

func (suite *PasswordServiceSuite) TestComparePassword_Failure() {
	password := "mySecureP@ssw0rd"
	hashedPassword, _ := suite.service.HashPassword(password)

	// Incorrect password should not match
	isMatch := suite.service.ComparePassword(hashedPassword, "wrongPassword")
	assert.False(suite.T(), isMatch)
}

func TestPasswordServiceSuite(t *testing.T) {
	suite.Run(t, new(PasswordServiceSuite))
}
