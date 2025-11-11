package service

import "github.com/payam1986128/go-fiber-sms-firewall/internal/repository"

type LimiterConditionService struct {
	repository *repository.LimiterConditionRepository
}

func NewLimiterConditionService(repository *repository.LimiterConditionRepository) *LimiterConditionService {
	return &LimiterConditionService{repository: repository}
}
