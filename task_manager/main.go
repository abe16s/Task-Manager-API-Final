package main

import (
    "fmt"
    "context"
    // "log"

    "github.com/abe16s/Go-Backend-Learning-path/task_manager/router"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/services"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/controllers"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://abeni:mongo123$@firstcluster.ae90a.mongodb.net/")

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		// log.Fatal(err)
        fmt.Errorf(err.Error())
	} else {
        fmt.Println("Connected to MongoDB!")
    }

	// // Check the connection
	// err = client.Ping(context.Background(), nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }


    taskService := services.TaskService{Client: client}
    taskController := controllers.TaskController{Service: taskService}
	// collection := client.Database("test").Collection("trainers")

    r := router.SetupRouter(&taskController)
    r.Run("localhost:8080")
}