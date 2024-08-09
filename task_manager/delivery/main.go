package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/delivery/controllers"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/delivery/router"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/usecases"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
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

	taskService := usecases.NewTaskService(client, dbName, "tasks")
	taskController := controllers.TaskController{Service: taskService}

	userService := usecases.NewUserService(client, dbName, "users")
	userController := controllers.UserController{Service: userService}
	
	r := router.SetupRouter(&taskController, &userController)
	r.Run("localhost:" + os.Getenv("SERVER_PORT"))
}