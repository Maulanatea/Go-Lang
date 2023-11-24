package midleware

import "github.com/gofiber/fiber/v2"

func Auth(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token != "secret" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauntenticated",
		})
	}

	return c.Next()
}

func PermissionCreate() {

}
