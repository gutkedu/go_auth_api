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
	if err := r.db.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

func (r *DBRepository) Update(
	ctx context.Context,
	roleID uuid.UUID,
	role *Role) error {
	if err := r.db.Model(&Role{}).Where("ID	= ?", roleID).Updates(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *DBRepository) Destroy(ctx context.Context, roleID uuid.UUID) error {
	if err := r.db.Delete(&Role{}, roleID).Error; err != nil {
		return err
	}
	return nil
}

func (r *DBRepository) Index(ctx context.Context) (*[]Role, error) {
	roles := []Role{}
	if err := r.db.Find(&roles).Error; err != nil {
		return &[]Role{}, err
	}
	return &roles, nil
}

func (r *DBRepository) Show(ctx context.Context, roleID uuid.UUID) (*Role, error) {
	role := Role{ID: roleID}
	if err := r.db.First(&role).Error; err != nil {
		return &Role{}, err
	}
	return &role, nil
}
