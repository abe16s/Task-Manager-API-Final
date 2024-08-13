package controllers

import (
	"net/http"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	Service usecases.UserServiceInterface
}

func (con *UserController) RegisterUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
		  validationErrors = errors
		}
	
		errorMessages := make(map[string]string)
		for _, e := range validationErrors {
	
		  field := e.Field()
		  switch field {
			case "Username":
				errorMessages["username"] = "username is required."
			case "Password":
				errorMessages["password"] = "password is required."
		  }
		}
	
		if len(errorMessages) == 0 {
			errorMessages["json"] = "Invalid JSON"
		}
		
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	newUser, err := con.Service.RegisterUser(&user)
	if err != nil && err.Error() == "username already exists" {
		c.IndentedJSON(http.StatusConflict, gin.H{"error": "username already exists"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}

func (con *UserController) Login(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := con.Service.LoginUser(user)
	if err != nil && err.Error() == "user not found" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil && err.Error() == "invalid credentials" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}

// promote user to admin
func (con *UserController) PromoteUser(c *gin.Context) {
	// get username from query parameter
	username := c.Query("username")
	err := con.Service.PromoteUser(username)
	if err != nil {
		if err.Error() =="username not found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}