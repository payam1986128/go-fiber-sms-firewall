package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/config"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) LoginHandler(c *fiber.Ctx) error {
	type cred struct {
		User string `json:"user"`
	}
	var rc cred
	if err := c.BodyParser(&rc); err != nil {
		return fiber.ErrBadRequest
	}
	// generate token (simple demo)
	token, err := config.GenerateJWT(rc.User)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{"token": token})
}
