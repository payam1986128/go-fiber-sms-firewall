package presentation

import "errors"

type (
	LoginUserRequest struct {
		Username *string `json:"username" validate:"required"`
	}

	RegisterUserRequest struct {
		Username *string `json:"username" validate:"required"`
	}

	VerificationResponse struct {
		AccessToken string `json:"accessToken"`
	}
)

func (req *LoginUserRequest) Validate() error {
	if req.Username == nil || *req.Username == "" {
		return errors.New("username is required")
	}
	return nil
}

func (req *RegisterUserRequest) Validate() error {
	if req.Username == nil || *req.Username == "" {
		return errors.New("username is required")
	}
	return nil
}
