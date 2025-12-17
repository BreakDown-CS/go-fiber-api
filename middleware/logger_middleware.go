package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {

		start := time.Now()

		duration := time.Since(start)

		log.Printf(
			"[REQUEST] Time : %v | %s %s ",
			duration,
			c.Method(),
			c.OriginalURL(),
		)

		return c.Next()
	}
}
