package midleware

import (
	"tutor-go-fiber/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauntenticated",
		})
	}
	_, err := utils.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauntenticated",
		})
	}

	// if token != "secret" {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"message": "unauntenticated",
	// 	})
	// }

	return c.Next()
}

func PermissionCreate() {

}
