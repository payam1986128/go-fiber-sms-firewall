package service

import (
	"os"
	"strings"
)

var countryCode = normalizeCountryCode()

func normalizePhoneNumber(phoneNumber *string) *string {
	if phoneNumber == nil || len(*phoneNumber) == 0 {
		return nil
	}
	if strings.HasPrefix(*phoneNumber, "00") {
		strings.Replace(*phoneNumber, "00", "+", 1)
	}
	if strings.HasPrefix(*phoneNumber, "0") {
		strings.Replace(*phoneNumber, "0", countryCode, 1)
	}
	return phoneNumber
}

func normalizeCountryCode() string {
	countryCode := os.Getenv("COUNTRY_CODE")
	strings.Replace(countryCode, "00", "+", 1)
	if !strings.HasPrefix(countryCode, "+") {
		countryCode = "+" + countryCode
	}
	return countryCode
}
