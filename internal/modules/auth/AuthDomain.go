package auth

import validation "github.com/go-ozzo/ozzo-validation"

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r AuthRequest) ValidateLoginInput() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}
