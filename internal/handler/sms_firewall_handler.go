package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/service"
)

type FirewallHandler struct {
	firewallService    *service.FirewallService
	transceiverService *service.TransceiverService
}

func NewFirewallHandler(firewallService *service.FirewallService, transceiverService *service.TransceiverService) *FirewallHandler {
	return &FirewallHandler{firewallService: firewallService, transceiverService: transceiverService}
}

func (handler *FirewallHandler) Receive(ctx *fiber.Ctx) error {
	var sms entity.Sms
	if err := ctx.BodyParser(&sms); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}
	receiverErr := handler.transceiverService.Receive(&sms)
	if receiverErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "transceiver receive error message: " + receiverErr.Error()})
	}
	protectErr := handler.firewallService.Protect(&sms)
	if protectErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "firewall protect error message: " + protectErr.Error()})
	}
	sendErr := handler.transceiverService.Send(&sms)
	if sendErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "transceiver send error message: " + sendErr.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"id": sms.ID.String()})
}
