package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) Register(ctx *fiber.Ctx) error {
	var code = ctx.Get("code")
	if len(code) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid code"})
	}
	var request presentation.RegisterUserRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}
	id, err := handler.service.RegisterUser(&request, code)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

func (handler *UserHandler) Login(ctx *fiber.Ctx) error {
	var code = ctx.Get("code")
	if len(code) == 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalid code"})
	}
	var request presentation.LoginUserRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}
	response, err := handler.service.LoginUser(request.Username, code)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(&response)
}
