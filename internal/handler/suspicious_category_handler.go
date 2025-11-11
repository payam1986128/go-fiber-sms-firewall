package handler

import "github.com/payam1986128/go-fiber-sms-firewall/internal/service"

type SuspiciousCategoryHandler struct {
	service *service.SuspiciousCategoryService
}

func NewSuspiciousCategoryHandler(service *service.SuspiciousCategoryService) *SuspiciousCategoryHandler {
	return &SuspiciousCategoryHandler{service: service}
}
