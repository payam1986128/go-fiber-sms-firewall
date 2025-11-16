package service

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
)

type LimiterConditionService struct {
	repository *repository.LimiterConditionRepository
}

func NewLimiterConditionService(repository *repository.LimiterConditionRepository) *LimiterConditionService {
	return &LimiterConditionService{repository: repository}
}

func (service *LimiterConditionService) GetLimiterConditions(request *presentation.LimiterConditionsSearchRequest) (*presentation.LimiterConditionsResponse, error) {
	return nil, nil
}

func (service *LimiterConditionService) GetLimiterCondition(id string) (*presentation.LimiterConditionResponse, error) {
	return nil, nil
}

func (service *LimiterConditionService) AddLimiterCondition(request *presentation.LimiterConditionRequest) (string, error) {
	return "", nil
}

func (service *LimiterConditionService) EditLimiterCondition(id string, request *presentation.LimiterConditionRequest) error {
	return nil
}

func (service *LimiterConditionService) ReviewLimiterCondition(request *presentation.LimiterConditionStateRequest) error {
	return nil
}

func (service *LimiterConditionService) DeleteLimiterCondition(id string) error {
	return nil
}

func (service *LimiterConditionService) GetCaughtSms(id string, request *presentation.SmsSearchRequest) (*presentation.SmsResponse, error) {
	return nil, nil
}
