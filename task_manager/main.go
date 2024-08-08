package main

import (
    "fmt"
    "context"
    "log"
    "github.com/abe16s/Go-Backend-Learning-path/task_manager/router"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/services"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/controllers"
)

func main() {
	clientOptions := options.Client().ApplyURI("")

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

    taskService := services.TaskService{Client: client}
    taskController := controllers.TaskController{Service: taskService}
    r := router.SetupRouter(&taskController)
    r.Run("localhost:8080")
}