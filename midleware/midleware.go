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
	// _, err := utils.VerifyToken(token)

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauntenticated",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden acces",
		})
	}

	// c.Locals("userInfo", claims)
	// c.Locals("role", claims["role"])

	return c.Next()
}

func PermissionCreate() {

}
