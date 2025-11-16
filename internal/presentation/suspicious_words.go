package presentation

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
