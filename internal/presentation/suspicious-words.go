package presentation

import "time"

type (
	SuspiciousCategoriesFilterRequest struct {
		Name string `json:"name"`
		Word string `json:"word"`
	}

	SuspiciousCategoriesResponse struct {
		Categories []SuspiciousCategoryDto `json:"categories"`
		Count      int                     `json:"count"`
	}

	SuspiciousCategoryDto struct {
		ID       string              `json:"id"`
		Name     string              `json:"name"`
		DateTime time.Time           `json:"dateTime"`
		Words    map[string]struct{} `json:"words"`
	}

	SuspiciousCategoryWordsRequest struct {
		Name  string              `json:"name"`
		Words map[string]struct{} `json:"words"`
	}
)
