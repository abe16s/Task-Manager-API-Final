package router

import (
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/controllers"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController) *gin.Engine {
    router := gin.Default()

    router.GET("/tasks", middleware.AuthMiddleware(false), taskController.GetTasks)
    router.GET("/tasks/:id", middleware.AuthMiddleware(false), taskController.GetTaskById)
    router.PUT("/tasks/:id", middleware.AuthMiddleware(true), taskController.UpdateTaskByID)
    router.DELETE("/tasks/:id", middleware.AuthMiddleware(true), taskController.DeleteTask)
    router.POST("/tasks", middleware.AuthMiddleware(true), taskController.AddTask)

	router.POST("/register", userController.RegisterUser)
	router.POST("/login", userController.Login)
    router.PATCH("/promote", middleware.AuthMiddleware(true), userController.PromoteUser)

    return router
}