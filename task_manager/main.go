package main

import (
    "github.com/abe16s/Go-Backend-Learning-path/task_manager/router"
)

func main() {
    r := router.SetupRouter()
    r.Run("localhost:8080")
}