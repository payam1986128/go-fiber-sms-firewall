package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/service"
)

type SuspiciousCategoryHandler struct {
	service *service.SuspiciousCategoryService
}

func NewSuspiciousCategoryHandler(service *service.SuspiciousCategoryService) *SuspiciousCategoryHandler {
	return &SuspiciousCategoryHandler{service: service}
}

func (handler *SuspiciousCategoryHandler) GetSuspiciousCategories(ctx *fiber.Ctx) error {
	var request presentation.SuspiciousCategoriesSearchRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}
	response, err := handler.service.GetSuspiciousCategories(&request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(&response)
}

func (handler *SuspiciousCategoryHandler) AddSuspiciousCategory(ctx *fiber.Ctx) error {
	var request presentation.SuspiciousCategoryWordsRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}
	if err := request.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&presentation.ValidationErrorDto{
			Message: err.Error(),
		})
	}
	id, err := handler.service.AddSuspiciousCategory(&request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

func (handler *SuspiciousCategoryHandler) EditSuspiciousCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := uuid.Validate(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}
	var request presentation.SuspiciousCategoryWordsRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}
	if err := request.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&presentation.ValidationErrorDto{
			Message: err.Error(),
		})
	}
	err := handler.service.EditSuspiciousCategory(id, &request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}

func (handler *SuspiciousCategoryHandler) DeleteSuspiciousCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := uuid.Validate(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}
	err := handler.service.DeleteSuspiciousCategory(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
