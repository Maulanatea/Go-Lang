package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const PathDefault = "./public/covers/"

func Singlefile(c *fiber.Ctx) error {
	//handle file
	file, err := c.FormFile("cover")
	if err != nil {
		log.Println("error file = ", err)
	}

	var fileName *string
	if file != nil {
		fileName = &file.Filename
		extenstionFile := filepath.Ext(*fileName)
		customFilename := fmt.Sprintf("gambar satu%s", extenstionFile)

		errSave := c.SaveFile(file, fmt.Sprintf("./public/covers/%s", customFilename))
		if errSave != nil {
			log.Println("fail to store file into public/cover")
		}
	} else {
		log.Println("tidak ada cover yg di updload")
	}

	if fileName != nil {
		c.Locals("filename", *fileName)
	} else {
		c.Locals("filename", nil)
	}

	return c.Next()
}

func MultipleFile(c *fiber.Ctx) error {
	form, errForm := c.MultipartForm()
	fmt.Println("keluarnya", form)
	if errForm != nil {
		log.Println("error read multipart form, error = ", errForm)
	}
	files := form.File["photos"]

	var filenames []string
	for i, file := range files {

		var fileName string
		if file != nil {
			fileName = fmt.Sprintf("%d-%s", i, file.Filename)
			errSave := c.SaveFile(file, fmt.Sprintf("./public/covers/%s", fileName))
			if errSave != nil {
				log.Println("fail to store file into public/cover")
			}
		} else {
			log.Println("tidak ada cover yg di updload")
		}

		if fileName != "" {
			filenames = append(filenames, fileName)
		}
	}

	c.Locals("filenames", filenames)

	return c.Next()

}

func RemoveFile(filename string, path ...string) error {
	if len(path) > 0 {
		err := os.Remove(path[0] + filename)
		if err != nil {
			log.Println("Failed To Remove File")
			return err
		}
	} else {
		err := os.Remove(PathDefault + filename)
		if err != nil {
			log.Println("Failed To Remove File")
			return err
		}
	}

	return nil
}
