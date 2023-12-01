package controllers

import (
	"fmt"
	"log"
	"tutor-go-fiber/database"
	"tutor-go-fiber/models/entity"
	"tutor-go-fiber/models/req"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookControllerAdd(c *fiber.Ctx) error {
	book := new(req.BookCreateReq)
	//pengecekan
	if err := c.BodyParser(book); err != nil {
		return err
	}
	//validasi sebelum user di buat
	validation := validator.New()
	if err := validation.Struct(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed input user",
			"error":   err.Error(),
		})
	}

	//handle file
	file, err := c.FormFile("cover")
	if err != nil {
		log.Println("error file = ", err)
	}

	fileName := file.Filename
	errSave := c.SaveFile(file, fmt.Sprintf("./public/covers/%s", fileName))
	if errSave != nil {
		log.Println("fail to store file into public/cover")
	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  fileName,
	}

	if err := database.DB.Create(&newBook).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed create new user",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Succes create new user",
		"data":    newBook,
	})
}
