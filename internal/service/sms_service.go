package service

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
)

type SmsService struct {
	repository *repository.SmsRepository
}

func NewSmsService(repository *repository.SmsRepository) *SmsService {
	return &SmsService{repository: repository}
}

func (service SmsService) GetSms(request *presentation.SmsSearchRequest) (*presentation.SmsResponse, error) {
	return nil, nil
}
