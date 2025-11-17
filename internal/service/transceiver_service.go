package service

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/config"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
	"time"
)

type TransceiverService struct {
	smsRepository      *repository.SmsRepository
	phoneNumberService *PhoneNumberService
	smscClient         *config.SmscClient
}

func NewTransceiverService(smsRepository *repository.SmsRepository, phoneNumberService *PhoneNumberService, smscClient *config.SmscClient) *TransceiverService {
	return &TransceiverService{
		smsRepository:      smsRepository,
		phoneNumberService: phoneNumberService,
		smscClient:         smscClient,
	}
}

func (service *TransceiverService) Receive(sms *entity.Sms) error {
	sms.ReceivedTime = time.Now()
	service.phoneNumberService.normalizePhoneNumber(&sms.Sender)
	service.phoneNumberService.normalizePhoneNumber(&sms.Receiver)
	return service.smsRepository.UpdateSms(sms)
}

func (service *TransceiverService) Send(sms *entity.Sms) error {
	if sms.Action == entity.DONT_SEND {
		return service.smsRepository.UpdateSms(sms)
	}
	request := presentation.SmsRequest{
		ID:       sms.ID,
		Sender:   sms.Sender,
		Receiver: sms.Receiver,
		Action:   sms.Action,
	}
	err := service.smscClient.SubmitSms(&request)
	if err != nil {
		return err
	}
	sms.SendTime = time.Now()
	return service.smsRepository.UpdateSms(sms)
}
