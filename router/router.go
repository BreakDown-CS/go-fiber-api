package router

import (
	"github.com/BreakDown-CS/go-fiber-api/handler"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/", handler.Health)
}
