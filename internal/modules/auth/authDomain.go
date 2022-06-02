package auth

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
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
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}

type AuthUseCase interface {
	AuthenticateUser(ctx context.Context, auth AuthRequest) (AuthResponse, error)
}
