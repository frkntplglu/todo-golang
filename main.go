package main

import (
	"log"

	"ToDo/database"
	"ToDo/handlers"
	"ToDo/repositories"
	"ToDo/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("cannot connected to db")
	}

	todoRepo := repositories.NewTodoRepository(db)
	err = todoRepo.Migration()
	if err != nil {
		log.Fatal(err)
	}
	todoService := services.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	app := fiber.New()
	app.Get("/api/todo", todoHandler.Get)
	app.Post("/api/todo", todoHandler.Create)
	app.Delete("/api/todo/:id", todoHandler.Delete)

	err = app.Listen(":8000")
	if err != nil {
		return
	}
}
