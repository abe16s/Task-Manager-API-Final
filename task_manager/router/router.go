package router

import (
    "github.com/gin-gonic/gin"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/controllers"
)

func SetupRouter(c *controllers.TaskController) *gin.Engine {
    router := gin.Default()

    router.GET("/tasks", c.GetTasks)
    router.GET("/tasks/:id", c.GetTaskById)
    router.PUT("/tasks/:id", c.UpdateTaskByID)
    router.DELETE("/tasks/:id", c.DeleteTask)
    router.POST("/tasks", c.AddTask)

    return router
}