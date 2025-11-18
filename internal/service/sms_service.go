package service

import (
	"fmt"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/mapper"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
	"time"
)

type SmsService struct {
	repository *repository.SmsRepository
}

func NewSmsService(repository *repository.SmsRepository) *SmsService {
	return &SmsService{repository: repository}
}

func (service SmsService) GetSms(request *presentation.SmsSearchRequest) (*presentation.SmsResponse, error) {
	where := ""
	currentTime := time.Now()
	if request.DateTo != nil {
		where += fmt.Sprintf("WHERE receivedTime < %d ", request.DateTo.UnixMilli())
	} else {
		where += fmt.Sprintf("WHERE receivedTime < %d ", currentTime.UnixMilli())
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
	count, err := service.repository.CountSmsByQuery(where)
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
	sms, err := service.repository.FindSmsByQuery(where)
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
