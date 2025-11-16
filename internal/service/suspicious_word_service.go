package service

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
)

type SuspiciousWordService struct {
	repository *repository.SuspiciousWordRepository
}

func NewSuspiciousWordService(repository *repository.SuspiciousWordRepository) *SuspiciousWordService {
	return &SuspiciousWordService{
		repository: repository,
	}
}

func (service *SuspiciousWordService) GetSuspiciousWords(request *presentation.SuspiciousWordsFilterRequest) (*presentation.SuspiciousWordsResponse, error) {
	return nil, nil
}

func (service *SuspiciousWordService) AddSuspiciousWords(request *presentation.SuspiciousWordsRequest) (string, error) {
	return "", nil
}

func (service *SuspiciousWordService) DeleteSuspiciousWords(id string) error {
	return nil
}
