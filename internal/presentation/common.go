package presentation

type (
	Pageable struct {
		Page     int `json:"page"`
		PageSize int `json:"pageSize"`
	}

	Sortable struct {
		Sort string `json:"sort"`
	}

	ValidationErrorDto struct {
		Message string            `json:"message"`
		Params  map[string]string `json:"params"`
	}
)
