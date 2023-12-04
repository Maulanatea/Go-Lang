package controllers

import (
	"log"
	"tutor-go-fiber/database"
	"tutor-go-fiber/models/entity"
	"tutor-go-fiber/models/req"
	"tutor-go-fiber/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserControllerShow(c *fiber.Ctx) error {
	var user []entity.User
	err := database.DB.Find(&user).Error
	if err != nil {
		log.Println(err)
	}
	//BISA JUGA MEMAKAI CODE DIBAWAH INI
	// result := database.DB.Find(&user)
	// if result.Error != nil {
	// 	log.Println(result.Error)
	// }
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
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	HashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	newUser.Password = HashedPassword

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

func UserControllerGetById(c *fiber.Ctx) error {
	var user []entity.User
	id := c.Params("id")
	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id tidak boleh kosong",
		})
		return nil
	}
	if err := database.DB.Where("id=?", id).First(&user).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "data user tidak ditemukan",
		})
		return nil
	}
	//ketika data di temukan
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil mengambil data user",
		"data":    user,
	})
}

func UserControllerUpdate(c *fiber.Ctx) error {
	userUP := new(req.UserUpdate)
	if err := c.BodyParser(userUP); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User

	id := c.Params("id")
	//cek available user
	if err := database.DB.Where("id=?", id).First(&user).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "update gagal",
		})
		return nil
	}
	//update user data
	if userUP.Name != "" {
		user.Name = userUP.Name
	}
	user.Email = userUP.Email
	err := database.DB.Save(&user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil mengupdate data user",
		"data":    user,
	})
}

func UserControllerDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entity.User

	//chek available user
	err := database.DB.Debug().First(&user, "id=?", id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "id tidak boleh kosong",
		})
	}

	//delete data user
	if err := database.DB.Debug().Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	return c.JSON(fiber.Map{
		"message": "data berhasil di delete",
	})
}
