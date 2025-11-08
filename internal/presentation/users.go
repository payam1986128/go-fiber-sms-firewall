package presentation

type (
	LoginUserRequest struct {
		Username string `json:"username"`
	}

	RegisterUserRequest struct {
		Username string `json:"username"`
	}

	VerificationResponse struct {
		AccessToken string `json:"accessToken"`
	}
)
