package main

import (
	"library_management/controllers"
	"library_management/models"
	"library_management/services"
)

func main() {
	lib := services.Library{Books: make(map[int]*models.Book), Members: make(map[int]*models.Member)}
	lib.Members[1] = &models.Member{Name: "Abenezer Seifu", ID: 1}
	controllers.StartConsole(&lib);
}
