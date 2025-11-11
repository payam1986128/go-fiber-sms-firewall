package handler

import "github.com/payam1986128/go-fiber-sms-firewall/internal/service"

type LimiterConditionHandler struct {
	service *service.LimiterConditionService
}

func NewLimiterConditionHandler(service *service.LimiterConditionService) *LimiterConditionHandler {
	return &LimiterConditionHandler{service: service}
}
