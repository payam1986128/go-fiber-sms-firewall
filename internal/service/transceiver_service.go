package service

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
)

type TransceiverService struct {
	smsRepository      *repository.SmsRepository
	phoneNumberService *PhoneNumberService
}

func NewTransceiverService(smsRepository *repository.SmsRepository, phoneNumberService *PhoneNumberService) *TransceiverService {
	return &TransceiverService{
		smsRepository:      smsRepository,
		phoneNumberService: phoneNumberService,
	}
}

func (s *TransceiverService) Send(sms *entity.Sms) error {
	return nil
}

func (s *TransceiverService) Receive(sms *entity.Sms) error {
	return nil
}
