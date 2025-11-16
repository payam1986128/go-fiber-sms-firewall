package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/service"
)

type SmsHandler struct {
	service *service.SmsService
}

func NewSmsHandler(service *service.SmsService) *SmsHandler {
	return &SmsHandler{service: service}
}

func (handler *SmsHandler) GetSms(ctx *fiber.Ctx) error {
	var request presentation.SmsFilterRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}
	response, err := handler.service.GetSms(&request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(&response)
}
