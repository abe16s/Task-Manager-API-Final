package router

import (
	"log"
	"os"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/delivery/controllers"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/infrastructure"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/usecases"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController) *gin.Engine {
    router := gin.Default()
    err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// get jwt secret from env
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	var jwtservice usecases.JwtServiceInterface = &infrastructure.JwtService{JwtSecret: jwtSecret}

    router.GET("/tasks", infrastructure.AuthMiddleware(jwtservice, false), taskController.GetTasks)
    router.GET("/tasks/:id", infrastructure.AuthMiddleware(jwtservice, false), taskController.GetTaskById)
    router.PUT("/tasks/:id", infrastructure.AuthMiddleware(jwtservice, true), taskController.UpdateTaskByID)
    router.DELETE("/tasks/:id", infrastructure.AuthMiddleware(jwtservice, true), taskController.DeleteTask)
    router.POST("/tasks", infrastructure.AuthMiddleware(jwtservice, true), taskController.AddTask)

	router.POST("/register", userController.RegisterUser)
	router.POST("/login", userController.Login)
    router.PATCH("/promote", infrastructure.AuthMiddleware(jwtservice, true), userController.PromoteUser)

    return router
}