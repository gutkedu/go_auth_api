package auth

import (
	"context"

	"github.com/gutkedu/golang_api/internal/modules/user"
)

type authUseCase struct {
	userRepository user.UserRepository
}

func NewAuthUseCase(r user.UserRepository) AuthUseCase {
	return &authUseCase{
		userRepository: r,
	}
}

func (s *authUseCase) findUserByEmail(
	ctx context.Context,
	email string) (*user.User, error) {
	return s.userRepository.FindByEmail(ctx, email)
}
