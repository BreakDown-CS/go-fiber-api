package handler

import (
	"strings"

	"github.com/BreakDown-CS/go-fiber-api/service"
	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {

	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	token, err := service.Login(req.Username, req.Password)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"access_token": token,
	})
}

func Logout(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	parts := strings.Split(auth, " ")
	token := parts[1]

	service.Logout(token)

	return c.JSON(fiber.Map{
		"message": "logout success",
	})
}
