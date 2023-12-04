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

func PhotoControllerCreate(c *fiber.Ctx) error {
	photo := new(req.PhotoCreateReq)
	//pengecekan
	if err := c.BodyParser(photo); err != nil {
		return err
	}
	//validasi sebelum user di buat
	validation := validator.New()
	if err := validation.Struct(photo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed input book",
			"error":   err.Error(),
		})
	}

	//validation ketika tidak ada cover/image
	filenamE := c.Locals("filenames")
	log.Println("filename = ", filenamE)
	if filenamE == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "image cover is required",
		})
	} else {
		fileNameData := filenamE.([]string)
		for _, filename := range fileNameData {
			newPhoto := entity.Photo{
				Image:      filename,
				CategoryId: photo.CategoryId,
			}

			if err := database.DB.Create(&newPhoto).Error; err != nil {
				log.Println("ada file yang gagal")

			}
		}

	}

	return c.JSON(fiber.Map{
		"message": "Succes create new photo",
	})
}

func PhotoControllerDelete(c *fiber.Ctx) error {
	var photos entity.Photo
	photoid := c.Params("id")

	//chek available photo
	err := database.DB.Debug().First(&photos, "id=?", photoid).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "id tidak boleh kosong",
		})
	}

	//mendelete file/photo yang ada disini
	errDelete := utils.RemoveFile(photos.Image)
	if errDelete != nil {
		log.Println("Fail to delete same file")
		return err
	}

	//delete data photo || ini mendelete untuk database nya
	if err := database.DB.Debug().Delete(&photos).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	return c.JSON(fiber.Map{
		"message": "photo berhasil di delete",
	})
}
