package roles

import (
	"context"

	"github.com/google/uuid"
)

type roleUseCase struct {
	roleRepository RoleRepository
}

func NewRoleUseCase(r RoleRepository) RoleUseCase {
	return &roleUseCase{
		roleRepository: r,
	}
}
func (s *roleUseCase) CreateRole(
	ctx context.Context,
	role *Role) error {
	return s.roleRepository.Store(ctx, role)
}

func (s *roleUseCase) GetRoles(ctx context.Context) (*[]Role, error) {
	return s.roleRepository.Index(ctx)
}

func (s *roleUseCase) GetRole(
	ctx context.Context,
	roleID uuid.UUID) (*Role, error) {
	return s.roleRepository.Show(ctx, roleID)
}

func (s *roleUseCase) UpdateRole(
	ctx context.Context,
	roleID uuid.UUID, role *Role) error {
	return s.roleRepository.Update(ctx, roleID, role)
}

func (s *roleUseCase) DeleteRole(
	ctx context.Context,
	roleID uuid.UUID) error {
	return s.roleRepository.Destroy(ctx, roleID)
}
