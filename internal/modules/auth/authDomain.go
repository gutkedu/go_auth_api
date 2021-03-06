package auth

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gutkedu/golang_api/internal/modules/user"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User  user.User `json:"user"`
	Token string    `json:"token"`
}

func (r AuthRequest) ValidateLoginInput() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required),
	)
}

type AuthUserUseCase interface {
	Execute(ctx context.Context, auth AuthRequest) (AuthResponse, error)
}
