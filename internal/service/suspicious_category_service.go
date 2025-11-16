package service

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
)

type SuspiciousCategoryService struct {
	repository *repository.SuspiciousCategoryRepository
}

func NewSuspiciousCategoryService(repository *repository.SuspiciousCategoryRepository) *SuspiciousCategoryService {
	return &SuspiciousCategoryService{repository: repository}
}

func (service SuspiciousCategoryService) GetSuspiciousCategories(request *presentation.SuspiciousCategoriesSearchRequest) (*presentation.SuspiciousCategoriesResponse, error) {
	return nil, nil
}

func (service SuspiciousCategoryService) AddSuspiciousCategory(request *presentation.SuspiciousCategoryWordsRequest) (string, error) {
	return "", nil
}

func (service SuspiciousCategoryService) EditSuspiciousCategory(id string, request *presentation.SuspiciousCategoryWordsRequest) error {
	return nil
}

func (service SuspiciousCategoryService) DeleteSuspiciousCategory(id string) error {
	return nil
}
