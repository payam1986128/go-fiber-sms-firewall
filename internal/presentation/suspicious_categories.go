package presentation

import (
	"errors"
	"time"
)

type (
	SuspiciousCategoriesSearchRequest struct {
		Pageable
		Sortable
		Name *string `json:"name"`
		Word *string `json:"word"`
	}

	SuspiciousCategoriesResponse struct {
		Categories []SuspiciousCategoryDto `json:"categories"`
		Count      int                     `json:"count"`
	}

	SuspiciousCategoryDto struct {
		ID       string    `json:"id"`
		Name     string    `json:"name"`
		DateTime time.Time `json:"dateTime"`
		Words    []string  `json:"words"`
	}

	SuspiciousCategoryWordsRequest struct {
		Name  *string  `json:"name" validate:"required"`
		Words []string `json:"words" validate:"required"`
	}
)

func (req *SuspiciousCategoryWordsRequest) Validate() error {
	if req.Name == nil {
		return errors.New("name is required")
	}
	if req.Words == nil || len(req.Words) == 0 {
		return errors.New("words is required")
	}
	return nil
}
