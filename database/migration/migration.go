package migration

import (
	"fmt"
	"tutor-go-fiber/database"
	"tutor-go-fiber/models/entity"
)

func RunMigrate() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Book{}, &entity.Category{}, &entity.Photo{})
	if err != nil {
		panic(err)
	}
	fmt.Println("success to migrate")
}
