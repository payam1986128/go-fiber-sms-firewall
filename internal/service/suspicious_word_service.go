package service

import "github.com/payam1986128/go-fiber-sms-firewall/internal/repository"

type SuspiciousWordService struct {
	repository *repository.SuspiciousWordRepository
}

func NewSuspiciousWordService(repository *repository.SuspiciousWordRepository) *SuspiciousWordService {
	return &SuspiciousWordService{
		repository: repository,
	}
}
