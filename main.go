package main

import (
	"github.com/BreakDown-CS/go-fiber-api/apperror"
	"github.com/BreakDown-CS/go-fiber-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if appErr, ok := err.(*apperror.AppError); ok {
				return c.Status(appErr.StatusCode).JSON(fiber.Map{
					"message": appErr.Message,
				})
			}

			return c.Status(500).JSON(fiber.Map{
				"message": "internal server error",
			})
		},
	})

	router.Register(app)

	app.Listen(":3000")
}
