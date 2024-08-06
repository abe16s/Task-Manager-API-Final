package router

import (
    "github.com/gin-gonic/gin"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/controllers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/tasks", controllers.GetTasks)
    router.GET("/tasks/:id", controllers.GetTaskById)
    router.PUT("/tasks/:id", controllers.UpdateTaskByID)
    router.DELETE("/tasks/:id", controllers.DeleteTask)
    router.POST("/tasks", controllers.AddTask)

    return router
}