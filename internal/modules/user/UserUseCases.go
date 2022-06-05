package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/gutkedu/golang_api/internal/modules/roles"
	"github.com/gutkedu/golang_api/internal/utils"
)

// Implementation of the repository in this service.
type userUseCase struct {
	userRepository UserRepository
	roleRepository roles.RoleRepository
}

// Create a new 'service' or 'use-case' for 'User' entity.
func NewUserUseCase(ur UserRepository, rr roles.RoleRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
		roleRepository: rr,
	}
}

func (s *userUseCase) GetUsers(ctx context.Context) (*[]User, error) {
	return s.userRepository.FindAll(ctx)
}

func (s *userUseCase) GetUser(ctx context.Context, userID uuid.UUID) (*User, error) {
	return s.userRepository.FindById(ctx, userID)
}

func (s *userUseCase) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return s.userRepository.FindByEmail(ctx, email)
}

func (s *userUseCase) CreateUser(ctx context.Context, user *User) error {
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash

	//TODO: assign new user to a role

	return s.userRepository.Create(ctx, user)
}

func (s *userUseCase) UpdateUser(ctx context.Context, userID uuid.UUID, user *User) error {
	user.UpdatedAt = time.Now()
	return s.userRepository.Update(ctx, userID, user)
}

func (s *userUseCase) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return s.userRepository.Delete(ctx, userID)
}
