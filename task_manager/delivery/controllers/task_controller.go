package controllers

import (
	"fmt"
	"net/http"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TaskController struct {
	Service usecases.TaskServiceInterface
}

func (con *TaskController) GetTasks(c *gin.Context) {
	tasks, err := con.Service.GetTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func (con *TaskController) GetTaskById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	task, err := con.Service.GetTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func (con *TaskController) UpdateTaskByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	var updatedTask domain.Task

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
	

	err = con.Service.UpdateTaskByID(id, updatedTask)

	if err != nil && err.Error() == "task not found" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (con *TaskController) DeleteTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	err = con.Service.DeleteTask(id)
	if err == nil {
		c.Status(http.StatusNoContent)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
}

func (con *TaskController) AddTask(c *gin.Context) {
	var newTask domain.Task
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

	task, err := con.Service.AddTask(newTask)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, task)
}
