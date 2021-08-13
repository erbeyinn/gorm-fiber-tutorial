package main

import (
	"erbeyinn/database"
	"erbeyinn/handlers"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := database.Connect()
	if err != nil {
		fmt.Println("Veritabanı hatası "+ err.Error())
	}

	err = database.AutoMigrater()
	if err != nil {
		fmt.Println("Veritabanı hatası "+ err.Error())
	}

	app.Get("/books", handlers.GetAllBooks)
	app.Post("/books",handlers.CreateNewBook)
	app.Get("/books/:id", handlers.GetBook)
	app.Delete("/books/:id", handlers.DeleteBook)

	app.Listen(":3000")
}