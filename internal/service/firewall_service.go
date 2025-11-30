package service

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
	"regexp"
	"strings"
	"time"
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
	activeLimiterConditions, err := service.limiterConditionService.GetActiveLimiterConditions()
	if err != nil {
		return err
	}
	for _, activeLimiterCondition := range activeLimiterConditions {
		if service.Evaluate(&activeLimiterCondition, sms) {
			service.Act(&activeLimiterCondition, sms)
			return service.smsRepository.UpdateSms(sms)
		}
	}
	service.Act(nil, sms)
	return service.smsRepository.UpdateSms(sms)
}

func (service *FirewallService) Act(condition *entity.LimiterCondition, sms *entity.Sms) {
	sms.Action = entity.SEND
	sms.EvaluatedTime = time.Now()
	if condition != nil {
		sms.AppliedFilterId = condition.ID
		sms.Action = condition.Action
	}
}

func (service *FirewallService) Evaluate(condition *entity.LimiterCondition, sms *entity.Sms) bool {
	return service.isConditionActive(condition) && service.isContentBlocked(condition, sms) ||
		service.isReceiverBlocked(condition, sms) || service.isSenderBlocked(condition, sms)
}

func (service *FirewallService) isConditionActive(condition *entity.LimiterCondition) bool {
	if condition.TimeLimits == nil || len(condition.TimeLimits) == 0 {
		return true
	}
	for _, timeLimit := range condition.TimeLimits {
		return (timeLimit.From == nil || time.Now().After(*timeLimit.From)) && time.Now().Before(timeLimit.To)
	}
	return false
}

func (service *FirewallService) isSenderBlocked(condition *entity.LimiterCondition, sms *entity.Sms) bool {
	senderFilter := condition.Filters.Sender
	if senderFilter == nil {
		return false
	}
	if senderFilter.Rate == nil && senderFilter.Senders != nil {
		for _, sender := range senderFilter.Senders {
			return sender == sms.Sender
		}
	}
	return service.isRateLimited(condition, sms)
}

func (service *FirewallService) isReceiverBlocked(condition *entity.LimiterCondition, sms *entity.Sms) bool {
	receivers := condition.Filters.Receivers
	if receivers != nil {
		for _, receiver := range receivers {
			return receiver == sms.Receiver
		}
	}
	return false
}

func (service *FirewallService) isRateLimited(condition *entity.LimiterCondition, sms *entity.Sms) bool {
	isSenderRateLimited := false
	if condition.Filters.Sender != nil && condition.Filters.Sender.Senders != nil {
		for _, sender := range condition.Filters.Sender.Senders {
			if sender == sms.Sender {
				isSenderRateLimited = true
				break
			}
		}
	}
	if isSenderRateLimited || condition.Filters.Sender == nil || condition.Filters.Sender.Senders == nil {
		smsCount, err := service.rateLimiterService.CountSms(condition, sms)
		if err != nil {
			return false
		}
		return smsCount >= condition.Filters.Sender.Rate.Threshold
	}
	return false
}

func (service *FirewallService) isContentBlocked(condition *entity.LimiterCondition, sms *entity.Sms) bool {
	keywordsFilter := condition.Filters.Keyword
	if keywordsFilter != nil {
		KeywordResult := false
		allKeywords := keywordsFilter.GetAllKeywords()
		for _, keyword := range allKeywords {
			KeywordResult = strings.Contains(sms.Message, keyword)
			break
		}
		regexResult := false
		for _, regex := range keywordsFilter.Regexes {
			regexResult, _ = regexp.MatchString(regex, sms.Message)
			break
		}
		return KeywordResult || regexResult
	}
	return false
}
