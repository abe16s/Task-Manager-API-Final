package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/delivery/controllers"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/delivery/router"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/repositories"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/usecases"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/infrastructure"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load("../.env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}

	dbName := "task-management"
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	var PasswordService usecases.PasswordServiceInterface = &infrastructure.PasswordService{}
	var JwtService usecases.JwtServiceInterface = &infrastructure.JwtService{JwtSecret: jwtSecret}

	var TaskRepository usecases.TaskRepoInterface = repositories.NewTaskRepository(client, dbName, "tasks")
	taskService := usecases.TaskService{TaskRepo: TaskRepository}
	taskController := controllers.TaskController{Service: &taskService}

	var UserRepository usecases.UserRepoInterface = repositories.NewUserRepository(client, dbName, "users")
	userService := usecases.UserService{UserRepo: UserRepository, PasswordService: PasswordService, JwtService: JwtService}
	userController := controllers.UserController{Service: &userService}
	
	r := router.SetupRouter(&taskController, &userController)
	r.Run("localhost:" + os.Getenv("SERVER_PORT"))
}
