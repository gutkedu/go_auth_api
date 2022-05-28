package user

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Implementation of the repository in this service.
type userService struct {
	userRepository UserRepository
}

// Create a new 'service' or 'use-case' for 'User' entity.
func NewUserService(r UserRepository) UserService {
	return &userService{
		userRepository: r,
	}
}

func (s *userService) GetUsers(ctx context.Context) (*[]User, error) {
	return s.userRepository.GetUsers(ctx)
}

func (s *userService) GetUser(ctx context.Context, userID uuid.UUID) (*User, error) {
	return s.userRepository.GetUser(ctx, userID)
}

func (s *userService) CreateUser(ctx context.Context, user *User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return s.userRepository.CreateUser(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, userID uuid.UUID, user *User) error {
	user.UpdatedAt = time.Now()
	return s.userRepository.UpdateUser(ctx, userID, user)
}

func (s *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return s.userRepository.DeleteUser(ctx, userID)
}
