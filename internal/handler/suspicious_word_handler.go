package handler

import "github.com/payam1986128/go-fiber-sms-firewall/internal/service"

type SuspiciousWordHandler struct {
	service *service.SuspiciousWordService
}

func NewSuspiciousWordHandler(service *service.SuspiciousWordService) *SuspiciousWordHandler {
	return &SuspiciousWordHandler{service: service}
}
