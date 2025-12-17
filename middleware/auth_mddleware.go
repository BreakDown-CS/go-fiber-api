package middleware

import "github.com/gofiber/fiber/v2"

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "missing authorization header",
			})
		}

		if authHeader != "Bearer valid-token" {
			return c.Status(401).JSON(fiber.Map{
				"message": "token expired or invalid",
			})
		}

		return c.Next()
	}
}
