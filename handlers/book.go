package handlers

import (
	"erbeyinn/database"
	"erbeyinn/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetBook(ctx *fiber.Ctx)error{
	var book models.Book

	id := ctx.Params("id")

	err := database.DB().Where("id = ?",id).First(&book).Error
	if err != nil {
		fmt.Println("Veritabanı hatası " + err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(book)

}

func CreateNewBook(ctx *fiber.Ctx) error{
	var book models.Book

	err := ctx.BodyParser(&book)
	if err != nil {
		fmt.Println("Veritabanı hatası" + err.Error())
	}

	err = database.DB().Create(&book).Error
	if err != nil {
		fmt.Println("Veritabanı hatası" + err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(book)
}


func GetAllBooks(ctx *fiber.Ctx)error{
	var books []models.Book
	err := database.DB().Find(&books).Error
	if err != nil {
		fmt.Println("Veritabanı hatası"+ err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(books)
}

func DeleteBook(ctx *fiber.Ctx) error{
	var book models.Book


	id := ctx.Params("id")

	err:= database.DB().Where("id = ?",id).Delete(&book).Error
	if err != nil {
		fmt.Println("Veritabanı hatası"+ err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(book)
}