package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/infrastructure"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/tests/mocks"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)


type AuthMiddlewareSuite struct {
	suite.Suite
	router *gin.Engine
	mockJwtService *mocks.JwtServiceInterface
}

func (suite *AuthMiddlewareSuite) SetupTest() {
	suite.router = gin.Default()
	suite.mockJwtService = new(mocks.JwtServiceInterface)
}

// Test AuthMiddleware with valid token

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_Success() {
	// Set up mock token validation
	token := &jwt.Token{}
	suite.mockJwtService.On("ValidateToken", "valid-token").Return(token, nil)
	suite.mockJwtService.On("ValidateAdmin", token).Return(true)

	// Create middleware with admin check enabled
	suite.router.Use(infrastructure.AuthMiddleware(suite.mockJwtService, true))
	suite.router.GET("/protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Access granted")
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Equal(suite.T(), "Access granted", rec.Body.String())
}

// Test AuthMiddleware with invalid token

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_InvalidToken() {
	suite.mockJwtService.On("ValidateToken", "invalid-token").Return(nil, errors.New("invalid JWT"))

	suite.router.Use(infrastructure.AuthMiddleware(suite.mockJwtService, false))
	suite.router.GET("/protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Access granted")
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusUnauthorized, rec.Code)
	assert.JSONEq(suite.T(), `{"error":"invalid JWT"}`, rec.Body.String())
}

// Test AuthMiddleware with missing header

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_MissingHeader() {
	suite.router.Use(infrastructure.AuthMiddleware(suite.mockJwtService, false))
	suite.router.GET("/protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Access granted")
	})

	req, _ := http.NewRequest("GET", "/protected", nil)

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusUnauthorized, rec.Code)
	assert.JSONEq(suite.T(), `{"error":"authorization header is required"}`, rec.Body.String())
}

// Test AuthMiddleware with invalid token bearer

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_InvalidHeader() {
	suite.router.Use(infrastructure.AuthMiddleware(suite.mockJwtService, false))
	suite.router.GET("/protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Access granted")
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "invalid-header-no-bearer")

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusUnauthorized, rec.Code)
	assert.JSONEq(suite.T(), `{"error": "invalid authorization header"}`, rec.Body.String())
}

// Test AuthMiddleware with admin check

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_AdminCheck_Failure() {
	token := &jwt.Token{}
	suite.mockJwtService.On("ValidateToken", "valid-token").Return(token, nil)
	suite.mockJwtService.On("ValidateAdmin", token).Return(false)

	suite.router.Use(infrastructure.AuthMiddleware(suite.mockJwtService, true))
	suite.router.GET("/admin", func(c *gin.Context) {
		c.String(http.StatusOK, "Admin access")
	})

	req, _ := http.NewRequest("GET", "/admin", nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusForbidden, rec.Code)
	assert.JSONEq(suite.T(), `{"error":"Forbidden"}`, rec.Body.String())
}

func TestAuthMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(AuthMiddlewareSuite))
}
