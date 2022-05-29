package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/gutkedu/golang_api/internal/modules/user/infra/gorm/entities"
	"gorm.io/gorm"
)

type mariaDBRepository struct {
	mariadb *gorm.DB
}

// Create a new repository with MariaDB as the driver.
func NewUserRepository(mariaDBConnection *gorm.DB) entities.UserRepository {
	return &mariaDBRepository{
		mariadb: mariaDBConnection,
	}
}

func (r *mariaDBRepository) GetUsers(ctx context.Context) (*[]entities.User, error) {
	users := []entities.User{}
	r.mariadb.Find(&users)
	return &users, nil
}

func (r *mariaDBRepository) GetUser(ctx context.Context, userID uuid.UUID) (*entities.User, error) {
	user := entities.User{ID: userID}
	r.mariadb.First(&user)
	return &user, nil
}

func (r *mariaDBRepository) CreateUser(ctx context.Context, user *entities.User) error {
	r.mariadb.Create(&user)
	return nil
}

func (r *mariaDBRepository) UpdateUser(ctx context.Context, userID uuid.UUID, user *entities.User) error {
	r.mariadb.Model(&entities.User{}).Where("id = ?", userID).Updates(user)
	return nil
}

func (r *mariaDBRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	r.mariadb.Delete(&entities.User{}, userID)
	return nil
}
