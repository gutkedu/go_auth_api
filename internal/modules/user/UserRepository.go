package user

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBRepository struct {
	db *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) UserRepository {
	return &DBRepository{
		db: dbConn,
	}
}

func (r *DBRepository) FindAll(ctx context.Context) (*[]User, error) {
	users := []User{}
	if err := r.db.Find(&users).Error; err != nil {
		return &[]User{}, err
	}
	return &users, nil
}

func (r *DBRepository) FindById(ctx context.Context, userID uuid.UUID) (*User, error) {
	user := User{ID: userID}
	if err := r.db.First(&user).Error; err != nil {
		return &User{}, err
	}
	return &user, nil
}

func (r *DBRepository) FindByEmail(ctx context.Context, userEmail string) (*User, error) {
	var user User
	if err := r.db.Where(&User{Email: userEmail}).Find(&user).Error; err != nil {
		return &User{}, err
	}
	return &user, nil
}

func (r *DBRepository) Create(ctx context.Context, user *User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *DBRepository) Update(ctx context.Context, userID uuid.UUID, user *User) error {
	if err := r.db.Model(&User{}).Where("id = ?", userID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *DBRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	if err := r.db.Delete(&User{}, userID).Error; err != nil {
		return err
	}
	return nil
}
