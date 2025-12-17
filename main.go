package main

import (
	"github.com/BreakDown-CS/go-fiber-api/middleware"
	"github.com/BreakDown-CS/go-fiber-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	router.Register(app)

	app.Listen(":3000")
}
