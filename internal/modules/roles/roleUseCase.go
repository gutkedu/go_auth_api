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

func (c *roleUseCase) GetRoles(ctx context.Context) (*[]Role, error) {
	return nil, nil
}

func (c *roleUseCase) GetRole(
	ctx context.Context,
	roleID uuid.UUID) (*Role, error) {
	return nil, nil
}

func (c *roleUseCase) CreateRole(
	ctx context.Context,
	role *Role) error {
	return nil
}

func (c *roleUseCase) UpdateRole(
	ctx context.Context,
	roleID uuid.UUID, role *Role) error {
	return nil
}

func (c *roleUseCase) DeleteRole(
	ctx context.Context,
	roleID uuid.UUID) error {
	return nil
}
