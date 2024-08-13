package tests

import (
	"testing"
	"time"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/infrastructure"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/usecases"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JwtServiceSuite struct {
	suite.Suite
	service usecases.JwtServiceInterface
	secret  []byte
}

func (suite *JwtServiceSuite) SetupTest() {
	suite.secret = []byte("test_secret")
	suite.service = &infrastructure.JwtService{JwtSecret: suite.secret}
}

// Test GenerateToken

func (suite *JwtServiceSuite) TestGenerateToken_Success() {
	token, err := suite.service.GenerateToken("testuser", true)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return suite.secret, nil
	})
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), parsedToken.Valid)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "testuser", claims["username"])
	assert.Equal(suite.T(), true, claims["is_admin"])
}


// Test ValidateToken

func (suite *JwtServiceSuite) TestValidateToken_Success() {
	token, _ := suite.service.GenerateToken("testuser", true)

	validatedToken, err := suite.service.ValidateToken(token)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), validatedToken)
	assert.True(suite.T(), validatedToken.Valid)

	claims, ok := validatedToken.Claims.(jwt.MapClaims)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "testuser", claims["username"])
	assert.Equal(suite.T(), true, claims["is_admin"])
}

func (suite *JwtServiceSuite) TestValidateToken_InvalidToken() {
	invalidToken := "invalid.token.string"

	_, err := suite.service.ValidateToken(invalidToken)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "invalid JWT", err.Error())
}

func (suite *JwtServiceSuite) TestValidateToken_ExpiredToken() {
	// Generate a token with a very short expiration time
	expirationTime := time.Now().Add(-1 * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "testuser",
		"is_admin": true,
		"exp":      expirationTime,
	})
	expiredToken, _ := token.SignedString(suite.secret)

	_, err := suite.service.ValidateToken(expiredToken)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "invalid JWT", err.Error())
}

// Test ValidateAdmin

func (suite *JwtServiceSuite) TestValidateAdmin_True() {
	token, _ := suite.service.GenerateToken("testuser", true)

	validatedToken, _ := suite.service.ValidateToken(token)
	isAdmin := suite.service.ValidateAdmin(validatedToken)
	assert.True(suite.T(), isAdmin)
}

func (suite *JwtServiceSuite) TestValidateAdmin_False() {
	token, _ := suite.service.GenerateToken("testuser", false)

	validatedToken, _ := suite.service.ValidateToken(token)
	isAdmin := suite.service.ValidateAdmin(validatedToken)
	assert.False(suite.T(), isAdmin)
}

func TestJwtServiceSuite(t *testing.T) {
	suite.Run(t, new(JwtServiceSuite))
}
