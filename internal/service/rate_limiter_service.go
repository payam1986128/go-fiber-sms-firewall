package service

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
	"time"
)

type RateLimiterService struct {
	*repository.SmsRepository
}

func NewRateLimiterService(smsRepository *repository.SmsRepository) *RateLimiterService {
	return &RateLimiterService{SmsRepository: smsRepository}
}

func (service *RateLimiterService) CountSms(condition *entity.LimiterCondition, sms *entity.Sms) (int, error) {
	intervalMinutes := condition.Filters.Sender.Rate.IntervalMinutes
	start := sms.ReceivedTime.Add(-time.Duration(intervalMinutes) * time.Minute)
	return service.SmsRepository.Count(sms.Sender, start, sms.ReceivedTime)
}
