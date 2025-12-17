package middleware

import (
	"github.com/BreakDown-CS/go-fiber-api/apperror"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {

	if appErr, ok := err.(*apperror.AppError); ok {
		return c.Status(appErr.StatusCode).JSON(fiber.Map{
			"message": appErr.Message,
		})
	}

	if fiberErr, ok := err.(*fiber.Error); ok {
		return c.Status(fiberErr.Code).JSON(fiber.Map{
			"message": fiberErr.Message,
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"message": "internal server error",
	})
}
