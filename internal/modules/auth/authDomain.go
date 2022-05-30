package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/gutkedu/golang_api/internal/modules/user"
)

type AuthUseCase interface {
	findUserByEmail(ctx context.Context, email string) (*user.User, error)
}

type LoginInput struct {
	email    string
	Password string
}

type UserData struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
