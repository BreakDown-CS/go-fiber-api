package handler

import (
	"strings"

	"github.com/BreakDown-CS/go-fiber-api/service"
	"github.com/BreakDown-CS/go-fiber-api/store"
	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func Login(c *fiber.Ctx) error {

	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	access, refresh, err := service.Login(req.Username, req.Password)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func Logout(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	token := strings.Split(auth, " ")[1]

	service.Logout(token)

	var req RefreshRequest

	_ = c.BodyParser(&req)
	store.Delete(req.RefreshToken)

	return c.JSON(fiber.Map{
		"message": "logout success",
	})
}

func Refresh(c *fiber.Ctx) error {
	var req RefreshRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	access, err := service.Refresh(req.RefreshToken)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"access_token": access,
	})
}
