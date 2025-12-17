package handler

import (
	"github.com/BreakDown-CS/go-fiber-api/service"
	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	result := service.ChackHealth()

	return c.SendString(result)
}

func GetUsers(c *fiber.Ctx) error {
	users, err := service.GetUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message":  "Internal Server Error",
			"Error : ": err.Error(),
		})
	}

	return c.Status(200).JSON(users)
}
