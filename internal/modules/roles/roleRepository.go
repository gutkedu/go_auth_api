package roles

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBRepository struct {
	db *gorm.DB
}

func NewRoleRepository(dbConn *gorm.DB) RoleRepository {
	return &DBRepository{
		db: dbConn,
	}
}

func (r *DBRepository) Store(ctx context.Context, role *Role) error {
	return nil
}

func (r *DBRepository) Update(ctx context.Context, roleID uuid.UUID) error {
	return nil
}

func (r *DBRepository) Destroy(ctx context.Context, roleID uuid.UUID) error {
	return nil
}

func (r *DBRepository) Index(ctx context.Context) (*[]Role, error) {
	return nil, nil
}

func (r *DBRepository) Show(ctx context.Context, roleID uuid.UUID) (role *Role, err error) {
	return nil, nil
}
