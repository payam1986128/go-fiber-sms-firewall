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

func (service *RateLimiterService) RateLimiter(sms entity.Sms, condition entity.LimiterCondition) (int, error) {
	intervalMinutes := condition.Filters.Sender.Rate.IntervalMinutes
	start := sms.ReceivedTime.Add(-time.Duration(intervalMinutes) * time.Minute)
	return service.SmsRepository.CountSms(sms.Sender, start, sms.ReceivedTime)
}
