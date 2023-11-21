package migration

import (
	"fmt"
	"tutor-go-fiber/database"
	"tutor-go-fiber/models/entity"
)

func RunMigrate() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
	fmt.Println("success to migrate")
}
