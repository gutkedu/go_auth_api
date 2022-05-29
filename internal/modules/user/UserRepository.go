package user

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type mariaDBRepository struct {
	mariadb *gorm.DB
}

// Create a new repository with MariaDB as the driver.
func NewUserRepository(mariaDBConnection *gorm.DB) UserRepository {
	return &mariaDBRepository{
		mariadb: mariaDBConnection,
	}
}

func (r *mariaDBRepository) GetUsers(ctx context.Context) (*[]User, error) {
	users := []User{}
	r.mariadb.Find(&users)
	return &users, nil
}

func (r *mariaDBRepository) GetUser(ctx context.Context, userID uuid.UUID) (*User, error) {
	user := User{ID: userID}
	r.mariadb.First(&user)
	return &user, nil
}

func (r *mariaDBRepository) CreateUser(ctx context.Context, user *User) error {
	r.mariadb.Create(&user)
	return nil
}

func (r *mariaDBRepository) UpdateUser(ctx context.Context, userID uuid.UUID, user *User) error {
	r.mariadb.Model(&User{}).Where("id = ?", userID).Updates(user)
	return nil
}

func (r *mariaDBRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	r.mariadb.Delete(&User{}, userID)
	return nil
}
