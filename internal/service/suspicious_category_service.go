package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/mapper"
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
	where := "WHERE 1=1 "
	if request.Name != nil && len(*request.Name) > 0 {
		where += fmt.Sprintf("and name LIKE '%%%s%%'", *request.Name)
	}
	if request.Word != nil && len(*request.Word) > 0 {
		where += fmt.Sprintf("and ANY w IN words SATISFIES w like '%%%s%%' END", *request.Word)
	}
	count, err := service.repository.CountByQuery(where)
	if err != nil {
		return nil, err
	}
	if request.Sort != nil && len(*request.Sort) > 0 {
		where += fmt.Sprintf(" ORDER BY %s", *request.Sort)
		if request.Dir != nil && len(*request.Dir) > 0 {
			where += fmt.Sprintf(" %s", *request.Dir)
		}
	}
	if request.PageSize != nil && *request.PageSize > 0 {
		where += fmt.Sprintf(" LIMIT %d", *request.PageSize)
	} else {
		where += fmt.Sprintf(" LIMIT 10")
	}
	if request.Page != nil && *request.Page > 0 {
		where += fmt.Sprintf(" OFFSET %d", *request.Page)
	}
	suspiciousCategories, err := service.repository.FindAllByQuery(where)
	if err != nil {
		return nil, err
	}
	dtos := mapper.ToSuspiciousCategoryDtos(suspiciousCategories)
	return &presentation.SuspiciousCategoriesResponse{
		Categories: dtos,
		Count:      count,
	}, nil
}

func (service SuspiciousCategoryService) AddSuspiciousCategory(request *presentation.SuspiciousCategoryWordsRequest) (string, error) {
	suspiciousCategory := mapper.ToSuspiciousCategory(*request)
	id, err := service.repository.Insert(&suspiciousCategory)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (service SuspiciousCategoryService) EditSuspiciousCategory(id string, request *presentation.SuspiciousCategoryWordsRequest) error {
	ID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	suspiciousCategory := mapper.ToSuspiciousCategory(*request)
	return service.repository.Update(ID, &suspiciousCategory)
}

func (service SuspiciousCategoryService) DeleteSuspiciousCategory(id string) error {
	ID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return service.repository.Delete(ID)
}
