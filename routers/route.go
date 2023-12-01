package routers

import (
	"tutor-go-fiber/controllers"
	"tutor-go-fiber/midleware"

	"github.com/gofiber/fiber/v2"
)

func RouterApp(c *fiber.App) {

	c.Post("/api/login", controllers.Login)                               //route untuk login
	c.Get("/api/showall", midleware.Auth, controllers.UserControllerShow) //route untuk menampilkan semua data
	c.Get("/api/showallById/:id", controllers.UserControllerGetById)      //route untuk menampilkan data menurut id
	c.Post("/api/create", controllers.UserControllerAdd)                  //route untuk menambahkan data
	c.Put("/api/updateById/:id", controllers.UserControllerUpdate)        //route untuk update data
	c.Delete("/api/delete/:id", controllers.UserControllerDelete)         //route untuk menghapus data

	c.Post("/api/bookCreate", controllers.BookControllerAdd)
}
