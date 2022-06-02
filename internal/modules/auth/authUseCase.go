package auth

import (
	"context"

	"github.com/gutkedu/golang_api/internal/modules/user"
)

// Implementation of the repository in this service.
type authUseCase struct {
	userRepository user.UserRepository
}

// Create a new 'service' or 'use-case' for 'User' entity.
func NewAuthUseCase(r user.UserRepository) AuthUseCase {
	return &authUseCase{
		userRepository: r,
	}
}

func (s *authUseCase) AuthenticateUser(
	ctx context.Context,
	auth AuthRequest) (AuthResponse, error) {
	user := user.User{}
	return AuthResponse{user, "token"}, nil
}
