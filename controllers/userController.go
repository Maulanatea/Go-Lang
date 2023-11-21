package controllers

import (
	"log"
	"tutor-go-fiber/database"
	"tutor-go-fiber/models/entity"
	"tutor-go-fiber/models/req"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserControllerShow(c *fiber.Ctx) error {
	var user []entity.User
	err := database.DB.Find(&user).Error
	if err != nil {
		log.Println(err)
	}
	return c.JSON(user)
}

func UserControllerAdd(c *fiber.Ctx) error {
	user := new(req.UserReq)
	//pengecekan
	if err := c.BodyParser(user); err != nil {
		return err
	}
	//validasi sebelum user di buat
	validation := validator.New()
	if err := validation.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed input user",
			"error":   err.Error(),
		})
	}

	newUser := entity.User{
		Name:  user.Name,
		Email: user.Email,
	}
	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed create new user",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Succes create new user",
		"data":    newUser,
	})
}
