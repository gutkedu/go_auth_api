package user

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid" json:"id"`
	Name      string    `gorm:"type:string;not null" json:"name"`
	Email     string    `gorm:"type:string;uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:string;not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	if err != nil {
		return err
	}
	return
}

func (r User) CreateUserValidation() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}

type UserRepository interface {
	FindAll(ctx context.Context) (*[]User, error)
	FindById(ctx context.Context, userID uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, userEmail string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, userID uuid.UUID, user *User) error
	Delete(ctx context.Context, userID uuid.UUID) error
}

type UserUseCase interface {
	GetUsers(ctx context.Context) (*[]User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, userID uuid.UUID, user *User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}
