package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/gutkedu/golang_api/internal/modules/user/infra/gorm/entities"
)

// Implementation of the repository in this service.
type userService struct {
	userRepository entities.UserRepository
}

// Create a new 'service' or 'use-case' for 'User' entity.
func NewUserUseCase(r entities.UserRepository) entities.UserService {
	return &userService{
		userRepository: r,
	}
}

func (s *userService) GetUsers(ctx context.Context) (*[]entities.User, error) {
	return s.userRepository.GetUsers(ctx)
}

func (s *userService) GetUser(ctx context.Context, userID uuid.UUID) (*entities.User, error) {
	return s.userRepository.GetUser(ctx, userID)
}

func (s *userService) CreateUser(ctx context.Context, user *entities.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return s.userRepository.CreateUser(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, userID uuid.UUID, user *entities.User) error {
	user.UpdatedAt = time.Now()
	return s.userRepository.UpdateUser(ctx, userID, user)
}

func (s *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return s.userRepository.DeleteUser(ctx, userID)
}
