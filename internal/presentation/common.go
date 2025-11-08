package presentation

type (
	ErrorDto struct {
		Message string `json:"message"`
	}

	Pageable struct {
		Page     int `json:"page"`
		PageSize int `json:"pageSize"`
	}

	Sortable struct {
		Sort string `json:"sort"`
	}

	SuccessfulCreationDto struct {
		ID string `json:"string"`
	}

	ValidationErrorDto struct {
		Message string            `json:"message"`
		Params  map[string]string `json:"params"`
	}
)
