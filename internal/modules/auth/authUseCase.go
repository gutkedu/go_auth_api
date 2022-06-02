package auth

import (
	"context"

	"github.com/gutkedu/golang_api/internal/modules/user"
	"github.com/gutkedu/golang_api/internal/utils"
)

// Implementation of the repository in this service.
type authUserUseCase struct {
	userRepository user.UserRepository
}

// Create a new 'service' or 'use-case' for 'User' entity.
func NewAuthUserUseCase(r user.UserRepository) AuthUserUseCase {
	return &authUserUseCase{
		userRepository: r,
	}
}

func (s *authUserUseCase) Execute(
	ctx context.Context,
	auth AuthRequest) (AuthResponse, error) {
	user := user.User{}

	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{user, token}, nil
}
