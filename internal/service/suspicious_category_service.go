package service

import "github.com/payam1986128/go-fiber-sms-firewall/internal/repository"

type SuspiciousCategoryService struct {
	repository *repository.SuspiciousCategoryRepository
}

func NewSuspiciousCategoryService(repository *repository.SuspiciousCategoryRepository) *SuspiciousCategoryService {
	return &SuspiciousCategoryService{repository: repository}
}
