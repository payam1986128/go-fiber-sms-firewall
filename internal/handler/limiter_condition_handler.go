package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/service"
)

type LimiterConditionHandler struct {
	service *service.LimiterConditionService
}

func NewLimiterConditionHandler(service *service.LimiterConditionService) *LimiterConditionHandler {
	return &LimiterConditionHandler{service: service}
}

func (handler *LimiterConditionHandler) GetLimiterConditions(ctx *fiber.Ctx) error {
	var request presentation.LimiterConditionsFilterRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}
	response, err := handler.service.GetLimiterConditions(&request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(&response)
}

func (handler *LimiterConditionHandler) GetLimiterCondition(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := uuid.Validate(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}
	response, err := handler.service.GetLimiterCondition(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "limiter condition not found"})
	}
	return ctx.Status(fiber.StatusOK).JSON(&response)
}

func (handler *LimiterConditionHandler) AddLimiterCondition(ctx *fiber.Ctx) error {
	var request presentation.LimiterConditionRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}
	id, err := handler.service.AddLimiterCondition(&request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

func (handler *LimiterConditionHandler) EditLimiterCondition(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := uuid.Validate(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}
	var request presentation.LimiterConditionRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}
	err := handler.service.EditLimiterCondition(id, &request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}

func (handler *LimiterConditionHandler) ReviewLimiterCondition(ctx *fiber.Ctx) error {
	var request presentation.LimiterConditionStateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}
	err := handler.service.ReviewLimiterCondition(&request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}

func (handler *LimiterConditionHandler) DeleteLimiterCondition(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := uuid.Validate(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}
	err := handler.service.DeleteLimiterCondition(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}

func (handler *LimiterConditionHandler) GetCaughtSms(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := uuid.Validate(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}
	var request presentation.SmsFilterRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}
	response, err := handler.service.GetCaughtSms(id, &request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(&response)
}
