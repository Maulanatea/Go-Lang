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
			"message": "failed input book",
			"error":   err.Error(),
		})
	}

	//validation ketika tidak ada cover/image
	var filenameString string
	filenamE := c.Locals("filename")
	log.Println("filename? = ", filenamE)
	if filenamE == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "image cover is required",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filenamE)

	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filenameString,
	}

	if err := database.DB.Create(&newBook).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed create new book",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Succes create new book",
		"data":    newBook,
	})
}
