package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/mapper"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
	"sync"
	"time"
)

var cache sync.Map

type LimiterConditionService struct {
	repository                   *repository.LimiterConditionRepository
	suspiciousCategoryRepository *repository.SuspiciousCategoryRepository
	smsRepository                *repository.SmsRepository
	phoneNumberService           *PhoneNumberService
}

func NewLimiterConditionService(repository *repository.LimiterConditionRepository, suspiciousCategoryRepository *repository.SuspiciousCategoryRepository, smsRepository *repository.SmsRepository, phoneNumberService *PhoneNumberService) *LimiterConditionService {
	return &LimiterConditionService{
		repository:                   repository,
		suspiciousCategoryRepository: suspiciousCategoryRepository,
		smsRepository:                smsRepository,
		phoneNumberService:           phoneNumberService,
	}
}

func (service *LimiterConditionService) GetActiveLimiterConditions() ([]entity.LimiterCondition, error) {
	cached, ok := cache.Load("ACTIVE_LIMITER_CONDITIONS")
	if ok {
		return cached.([]entity.LimiterCondition), nil
	}
	limiterConditions, err := service.repository.FindActiveLimiterConditions()
	if err != nil {
		return nil, err
	}
	cache.Store("ACTIVE_LIMITER_CONDITIONS", limiterConditions)
	return limiterConditions, nil
}

func (service *LimiterConditionService) GetLimiterConditions(request *presentation.LimiterConditionsSearchRequest) (*presentation.LimiterConditionsResponse, error) {
	where := "WHERE 1=1"
	if request.Filter != nil {
		where += fmt.Sprintf("AND name LIKE '%%%s%%' ", *request.Filter)
	}
	if request.State != nil {
		where += fmt.Sprintf("AND active LIKE '%%%t%%' ", *request.State)
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
	limiterConditions, err := service.repository.FindAllByQuery(where)
	if err != nil {
		return nil, err
	}
	dtos := mapper.ToBriefLimiterConditionDtos(limiterConditions)
	return &presentation.LimiterConditionsResponse{
		LimiterConditions: dtos,
		Count:             count,
		SearchTime:        time.Now(),
	}, nil
}

func (service *LimiterConditionService) GetLimiterCondition(id string) (*presentation.LimiterConditionResponse, error) {
	ID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	found, err := service.repository.GetByID(ID)
	if err != nil {
		return nil, err
	}
	result := mapper.ToLimiterConditionResponse(found)
	result.CaughtSms, _ = service.smsRepository.CountByAppliedFilterId(ID)
	return &result, nil
}

func (service *LimiterConditionService) EditContainingLimitConditions(categoryId uuid.UUID) error {
	where := fmt.Sprintf("where ANY c IN filters.keyword.Categories SATISFIES c = '%%%s%%' END", categoryId.String())
	conditions, err := service.repository.FindAllByQuery(where)
	if err != nil {
		return err
	}
	for _, condition := range conditions {
		service.updateLimiterConditionCategories(&condition)
		err := service.repository.Update(condition.ID, &condition)
		if err != nil {
			return err
		}
	}
	cache.Delete("ACTIVE_LIMITER_CONDITIONS")
	return nil
}

func (service *LimiterConditionService) DeleteFromContainingLimitConditions(categoryId uuid.UUID) error {
	where := fmt.Sprintf("where ANY c IN filters.keyword.Categories SATISFIES c = '%%%s%%' END", categoryId.String())
	conditions, err := service.repository.FindAllByQuery(where)
	if err != nil {
		return err
	}
	for _, condition := range conditions {
		categoryIds := condition.Filters.Keyword.Categories
		foundIndex := 0
		for i, id := range categoryIds {
			if id == categoryId {
				foundIndex = i
			}
		}
		categoryIds = append(categoryIds[:foundIndex], categoryIds[foundIndex+1:]...)
		service.updateLimiterConditionCategories(&condition)
		err := service.repository.Update(condition.ID, &condition)
		if err != nil {
			return err
		}
	}
	cache.Delete("ACTIVE_LIMITER_CONDITIONS")
	return nil
}

func (service *LimiterConditionService) AddLimiterCondition(request *presentation.LimiterConditionRequest) (string, error) {
	service.normalizePhoneNumbers(request)
	limiterCondition := entity.LimiterCondition{}
	mapper.ToLimiterCondition(*request, &limiterCondition)
	service.updateLimiterConditionCategories(&limiterCondition)
	saved, err := service.repository.Insert(&limiterCondition)
	if err != nil {
		return "", err
	}
	cache.Delete("ACTIVE_LIMITER_CONDITIONS")
	return saved.String(), nil
}

func (service *LimiterConditionService) EditLimiterCondition(id string, request *presentation.LimiterConditionRequest) error {
	ID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	service.normalizePhoneNumbers(request)
	foundLimiterCondition, err := service.repository.GetByID(ID)
	if err != nil {
		return err
	}
	mapper.ToLimiterCondition(*request, foundLimiterCondition)
	service.updateLimiterConditionCategories(foundLimiterCondition)
	err = service.repository.Update(ID, foundLimiterCondition)
	if err != nil {
		return err
	}
	cache.Delete("ACTIVE_LIMITER_CONDITIONS")
	return nil
}

func (service *LimiterConditionService) updateLimiterConditionCategories(limiterCondition *entity.LimiterCondition) {
	keywordsFilter := limiterCondition.Filters.Keyword
	var categories []entity.SuspiciousCategory
	var err error
	if keywordsFilter != nil && keywordsFilter.Categories != nil {
		categories, err = service.suspiciousCategoryRepository.FindAllByIds(keywordsFilter.Categories)
		if err != nil {
			return
		}
		keywordsFilter.CategoryKeywords = make([]string, 0)
		for _, category := range categories {
			keywordsFilter.CategoryKeywords = append(keywordsFilter.CategoryKeywords, category.Words...)
		}
	}
}

func (service *LimiterConditionService) normalizePhoneNumbers(request *presentation.LimiterConditionRequest) {
	if request.Filters.Sender != nil {
		for _, sender := range request.Filters.Sender.Senders {
			service.phoneNumberService.normalizePhoneNumber(&sender)
		}
	}
	if request.Filters.Receivers != nil {
		for _, receiver := range request.Filters.Receivers {
			service.phoneNumberService.normalizePhoneNumber(&receiver)
		}
	}
}

func (service *LimiterConditionService) ReviewLimiterCondition(request *presentation.LimiterConditionStateRequest) error {
	ids := make([]uuid.UUID, len(request.IDs))
	for i, id := range request.IDs {
		ids[i], _ = uuid.Parse(id)
	}
	limiterConditions, err := service.repository.FindAllByIds(ids)
	if err != nil {
		return err
	}
	for _, limiterCondition := range limiterConditions {
		limiterCondition.Active = *request.State
		err := service.repository.Update(limiterCondition.ID, &limiterCondition)
		if err != nil {
			return err
		}
	}
	cache.Delete("ACTIVE_LIMITER_CONDITIONS")
	return nil
}

func (service *LimiterConditionService) DeleteLimiterCondition(id string) error {
	ID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	cache.Delete("ACTIVE_LIMITER_CONDITIONS")
	return service.repository.Delete(ID)
}

func (service *LimiterConditionService) GetCaughtSms(id string, request *presentation.SmsSearchRequest) (*presentation.SmsResponse, error) {
	where := fmt.Sprintf("WHERE appliedFilterId = %s", id)
	currentTime := time.Now()
	if request.DateTo != nil {
		where += fmt.Sprintf("and receivedTime < %d ", request.DateTo.UnixMilli())
	} else {
		where += fmt.Sprintf("and receivedTime < %d ", currentTime.UnixMilli())
	}
	if request.DateFrom != nil {
		where += fmt.Sprintf("and receivedTime > %d", request.DateFrom.UnixMilli())
	}
	if request.Sender != nil && len(*request.Sender) > 0 {
		where += fmt.Sprintf(" and sender = %s", *request.Sender)
	}
	if request.Receiver != nil && len(*request.Receiver) > 0 {
		where += fmt.Sprintf(" and receiver = %s", *request.Receiver)
	}
	if request.Action != nil && len(*request.Action) > 0 {
		where += fmt.Sprintf(" and action = %s", *request.Action)
	}
	count, err := service.smsRepository.CountSmsByQuery(where)
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
	sms, err := service.smsRepository.FindAllByQuery(where)
	if err != nil {
		return nil, err
	}
	dtos := mapper.ToBriefSmsDtos(sms)
	return &presentation.SmsResponse{
		Sms:        dtos,
		Count:      count,
		SearchTime: currentTime,
	}, nil
}
