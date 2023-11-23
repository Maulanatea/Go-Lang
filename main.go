package main

import (
	"tutor-go-fiber/database"
	"tutor-go-fiber/database/migration"
	"tutor-go-fiber/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()   //initial database
	migration.RunMigrate() //run migration
	app := fiber.New()

	routers.RouterApp(app)

	app.Listen(":8000")
}
