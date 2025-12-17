package handler

import (
	"github.com/BreakDown-CS/go-fiber-api/service"
	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	result := service.ChackHealth()
	return c.SendString(result)
}
