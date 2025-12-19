package router

import (
	"github.com/BreakDown-CS/go-fiber-api/handler"
	"github.com/BreakDown-CS/go-fiber-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {

	app.Use(middleware.Logger())

	app.Get("/", handler.Health)

	app.Post("/login", handler.Login)

	app.Post("/refresh", handler.Refresh)

	protected := app.Group("/api", middleware.Auth())

	protected.Get("/users", handler.GetUsers)
	protected.Post("/logout", handler.Logout)
}
