package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/service"
)

type SuspiciousWordHandler struct {
	service *service.SuspiciousWordService
}

func NewSuspiciousWordHandler(service *service.SuspiciousWordService) *SuspiciousWordHandler {
	return &SuspiciousWordHandler{service: service}
}

func (handler *SuspiciousWordHandler) GetSuspiciousWords(ctx *fiber.Ctx) error {
	var request presentation.SuspiciousWordsFilterRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}
	response, err := handler.service.GetSuspiciousWords(&request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(&response)
}

func (handler *SuspiciousWordHandler) AddSuspiciousWords(ctx *fiber.Ctx) error {
	var request presentation.SuspiciousWordsRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}
	id, err := handler.service.AddSuspiciousWords(&request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

func (handler *SuspiciousWordHandler) DeleteSuspiciousWords(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := uuid.Validate(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}
	err := handler.service.DeleteSuspiciousWords(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
