package service

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
)

type FirewallService struct {
	smsRepository           *repository.SmsRepository
	rateLimiterService      *RateLimiterService
	limiterConditionService *LimiterConditionService
}

func NewFirewallService(smsRepository *repository.SmsRepository, rateLimiterService *RateLimiterService, limiterConditionService *LimiterConditionService) *FirewallService {
	return &FirewallService{smsRepository: smsRepository, rateLimiterService: rateLimiterService, limiterConditionService: limiterConditionService}
}

func (service *FirewallService) Protect(sms *entity.Sms) error {
	return nil
}
