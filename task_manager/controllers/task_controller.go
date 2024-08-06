package controllers

import (
	"fmt"
	"net/http"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/models"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetTasks(c *gin.Context) {
	tasks := services.GetTasks()
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := services.GetTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func UpdateTaskByID(c *gin.Context) {
	id := c.Param("id")

	var updatedTask models.Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {

		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
		  validationErrors = errors
		}
	
		errorMessages := make(map[string]string)
		for _, e := range validationErrors {
	
		  field := e.Field()
		  fmt.Println(field, "this is field")
		  switch field {
		  case "Title":
			errorMessages["title"] = "Title is required."
		  case "Description":
			errorMessages["description"] = "Description is required."
	
		  case "Status":
			errorMessages["status"] = "Status is required."
	
		  case "DueDate":
			errorMessages["due_date"] = "DueDate is required."
		  }
		}
	
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	  }
	

	task, err := services.UpdateTaskByID(id, updatedTask)

	if err != nil && err.Error() == "task not found" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	deletedTask, err := services.DeleteTask(id)
	if err == nil {
		c.IndentedJSON(http.StatusAccepted, deletedTask)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
}

func AddTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {

		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
		  validationErrors = errors
		}
	
		errorMessages := make(map[string]string)
		for _, e := range validationErrors {
	
		  field := e.Field()
		  fmt.Println(field, "this is field")
		  switch field {
		  case "Title":
			errorMessages["title"] = "Title is required."
		  case "Description":
			errorMessages["description"] = "Description is required."
	
		  case "Status":
			errorMessages["status"] = "Status is required."
	
		  case "DueDate":
			errorMessages["due_date"] = "DueDate is required."
		  }
		}
	
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	  }
	
	baseURL := fmt.Sprintf("http://%s", c.Request.Host)

	resourceLocation := fmt.Sprintf("%s%s/%s", baseURL, c.Request.URL.Path, newTask.ID)
	c.Header("Location", resourceLocation)

	task := services.AddTask(newTask)
	c.IndentedJSON(http.StatusCreated, task)
}
