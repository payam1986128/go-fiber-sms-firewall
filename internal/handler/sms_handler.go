package handler

import "github.com/payam1986128/go-fiber-sms-firewall/internal/service"

type SmsHandler struct {
	service *service.SmsService
}

func NewSmsHandler(service *service.SmsService) *SmsHandler {
	return &SmsHandler{service: service}
}
