package controllers

import (
	"log"
	"time"
	"tutor-go-fiber/database"
	"tutor-go-fiber/models/entity"
	"tutor-go-fiber/models/req"
	"tutor-go-fiber/utils"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(req.LoginReq)

	//pengecekan
	if err := c.BodyParser(loginRequest); err != nil {
		return err
	}

	//validasi request
	validation := validator.New()
	if err := validation.Struct(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	//check available user
	var user entity.User
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "user ini tidak ada",
		})
	}

	//check validation password
	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credencial",
		})
	}

	// GENERATE JWT
	claimss := jwt.MapClaims{}
	claimss["name"] = user.Name
	claimss["email"] = user.Email
	claimss["exp"] = time.Now().Add(time.Minute * 2).Unix()

	if user.Email == "samsudin@gmail.com" {
		claimss["role"] = "admin"
	} else {
		claimss["role"] = "user"

	}

	token, err := utils.GenerateToken(&claimss)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token tidak valid",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
