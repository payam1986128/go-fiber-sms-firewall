package presentation

import "time"

type (
	SuspiciousWordDto struct {
		ID       string    `json:"id"`
		Word     string    `json:"word"`
		DateTime time.Time `json:"dateTime"`
	}

	SuspiciousWordsFilterRequest struct {
		Filter string `json:"filter"`
	}

	SuspiciousWordsRequest struct {
		Words map[string]struct{} `json:"words"`
	}

	SuspiciousWordsResponse struct {
		Words []SuspiciousWordDto `json:"words"`
		Count int                 `json:"count"`
	}
)
