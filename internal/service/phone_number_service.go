package service

import (
	"os"
	"strings"
)

type PhoneNumberService struct {
	countryCode string
}

func NewPhoneNumberService() *PhoneNumberService {
	service := &PhoneNumberService{
		countryCode: os.Getenv("COUNTRY_CODE"),
	}
	service.normalizeCountryCode()
	return service
}

func (service *PhoneNumberService) normalizePhoneNumber(phoneNumber *string) *string {
	if phoneNumber == nil || len(*phoneNumber) == 0 {
		return nil
	}
	if strings.HasPrefix(*phoneNumber, "00") {
		strings.Replace(*phoneNumber, "00", "+", 1)
	}
	if strings.HasPrefix(*phoneNumber, "0") {
		strings.Replace(*phoneNumber, "0", service.countryCode, 1)
	}
	return phoneNumber
}

func (service *PhoneNumberService) normalizeCountryCode() {
	strings.Replace(service.countryCode, "00", "+", 1)
	if !strings.HasPrefix(service.countryCode, "+") {
		service.countryCode = "+" + service.countryCode
	}
}
