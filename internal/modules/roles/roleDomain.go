package roles

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID `gorm:"type:uuid" json:"id"`
	Privilege   string    `gorm:"type:string" json:"privilege"`
	Description string    `gorm:"type:string" json:"description"`
	CreatedAt   time.Time `gorm:"autoUpdateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoCreateTime" json:"updated_at"`
}

func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
	role.ID = uuid.New()
	if err != nil {
		return err
	}
	return
}

func (r Role) CreateRoleValidation() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Privilege, validation.Required),
		validation.Field(&r.Description, validation.Required),
	)
}

type RoleRepository interface {
	Store(ctx context.Context, role *Role) error
	Update(ctx context.Context, roleID uuid.UUID) error
	Destroy(ctx context.Context, roleID uuid.UUID) error
	Index(ctx context.Context) (*[]Role, error)
	Show(ctx context.Context, roleID uuid.UUID) (role *Role, err error)
}

type RoleUseCase interface {
	GetRoles(ctx context.Context) (*[]Role, error)
	GetRole(ctx context.Context, roleID uuid.UUID) (*Role, error)
	DeleteRole(ctx context.Context, roleID uuid.UUID) error
	UpdateRole(ctx context.Context, roleID uuid.UUID, role *Role) error
	CreateRole(ctx context.Context, role *Role) error
}
