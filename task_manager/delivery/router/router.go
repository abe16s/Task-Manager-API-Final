package router

import (
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/delivery/controllers"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/infrastructure"
	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController) *gin.Engine {
    router := gin.Default()

    router.GET("/tasks", infrastructure.AuthMiddleware(false), taskController.GetTasks)
    router.GET("/tasks/:id", infrastructure.AuthMiddleware(false), taskController.GetTaskById)
    router.PUT("/tasks/:id", infrastructure.AuthMiddleware(true), taskController.UpdateTaskByID)
    router.DELETE("/tasks/:id", infrastructure.AuthMiddleware(true), taskController.DeleteTask)
    router.POST("/tasks", infrastructure.AuthMiddleware(true), taskController.AddTask)

	router.POST("/register", userController.RegisterUser)
	router.POST("/login", userController.Login)
    router.PATCH("/promote", infrastructure.AuthMiddleware(true), userController.PromoteUser)

    return router
}