package routers

import (
	"tutor-go-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouterApp(c *fiber.App) {
	c.Get("/api/showall", controllers.UserControllerShow)
	c.Get("/api/showallById/:id", controllers.UserControllerGetById)
	c.Post("/api/create", controllers.UserControllerAdd)
}
