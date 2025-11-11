package presentation

import "time"

type (
	SuspiciousCategoriesFilterRequest struct {
		Pageable
		Sortable
		Name string `json:"name"`
		Word string `json:"word"`
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
		Name  string   `json:"name"`
		Words []string `json:"words"`
	}
)
