package controllers

import (
	"tutor-go-fiber/database"
	"tutor-go-fiber/models/entity"
	"tutor-go-fiber/models/req"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(req.LoginReq)

	//pengecekan
	if err := c.BodyParser(loginRequest); err != nil {
		return err
	}

	//validasi sebelum user di buat
	validation := validator.New()
	if err := validation.Struct(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	var user entity.User
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "user ini tidak ada",
		})
	}

	return c.JSON(fiber.Map{
		"token": "secret",
	})
}
