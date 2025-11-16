package presentation

import "errors"

type (
	SuspiciousWordDto struct {
		ID       string `json:"id"`
		Word     string `json:"word"`
		DateTime string `json:"dateTime"`
	}

	SuspiciousWordsFilterRequest struct {
		Pageable
		Sortable
		Filter string `json:"filter"`
	}

	SuspiciousWordsRequest struct {
		Words []string `json:"words"`
	}

	SuspiciousWordsResponse struct {
		Words []SuspiciousWordDto `json:"words"`
		Count int                 `json:"count"`
	}
)

func (req *SuspiciousWordsRequest) Validate() error {
	if req.Words == nil || len(req.Words) == 0 {
		return errors.New("words is required")
	}
	return nil
}
