package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Implementation of the repository in this service.
type userUseCase struct {
	userRepository UserRepository
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Create a new 'service' or 'use-case' for 'User' entity.
func NewUserUseCase(r UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: r,
	}
}

func (s *userUseCase) GetUsers(ctx context.Context) (*[]User, error) {
	return s.userRepository.GetUsers(ctx)
}

func (s *userUseCase) GetUser(ctx context.Context, userID uuid.UUID) (*User, error) {
	return s.userRepository.GetUser(ctx, userID)
}

func (s *userUseCase) CreateUser(ctx context.Context, user *User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hash, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash

	return s.userRepository.CreateUser(ctx, user)
}

func (s *userUseCase) UpdateUser(ctx context.Context, userID uuid.UUID, user *User) error {
	user.UpdatedAt = time.Now()
	return s.userRepository.UpdateUser(ctx, userID, user)
}

func (s *userUseCase) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return s.userRepository.DeleteUser(ctx, userID)
}