package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/mapper"
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
<<<<<<< HEAD
	where := ""
	if request.Filter != nil && len(*request.Filter) > 0 {
		where += fmt.Sprintf("WHERE word LIKE '%%%s%%'", *request.Filter)
	}
=======
	where := fmt.Sprintf("WHERE word LIKE '%%%s%%'", request.Filter)
>>>>>>> origin/feature/initial
	count, err := service.repository.CountSuspiciousWordsByQuery(where)
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
	suspiciousWords, err := service.repository.FindSuspiciousWordsByQuery(where)
	if err != nil {
		return nil, err
	}
	dtos := mapper.ToSuspiciousWordDtos(suspiciousWords)
	return &presentation.SuspiciousWordsResponse{
		Words: dtos,
		Count: count,
	}, nil
}

func (service *SuspiciousWordService) AddSuspiciousWords(request *presentation.SuspiciousWordsRequest) error {
	suspiciousWords := make([]entity.SuspiciousWord, len(request.Words))
	for _, word := range request.Words {
		suspiciousWords = append(suspiciousWords, entity.SuspiciousWord{
			Word: word,
		})
	}
	return service.repository.AddSuspiciousWords(suspiciousWords)
}

func (service *SuspiciousWordService) DeleteSuspiciousWords(id string) error {
	ID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return service.repository.DeleteSuspiciousWord(ID)
}
