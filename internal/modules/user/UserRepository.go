package user

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBRepository struct {
	db *gorm.DB
}

// Create a new repository with MariaDB as the driver.
func NewUserRepository(dbConn *gorm.DB) UserRepository {
	return &DBRepository{
		db: dbConn,
	}
}

func (r *DBRepository) FindAll(ctx context.Context) (*[]User, error) {
	users := []User{}
	r.db.Find(&users)
	return &users, nil
}

func (r *DBRepository) FindById(ctx context.Context, userID uuid.UUID) (*User, error) {
	user := User{ID: userID}
	r.db.First(&user)
	return &user, nil
}

/*
func (r *DBRepository) FindByEmail(ctx context.Context, userEmail string) (*User, error) {
	var user User
	if err := r.db.Where(&User{Email: userEmail}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
*/

func (r *DBRepository) FindByEmail(ctx context.Context, userEmail string) (*User, error) {
	//db.Where("name = ?", "jinzhu").First(&user)
	var user User
	r.db.Where(&User{Email: userEmail}).Find(&user)
	return &user, nil
}

func (r *DBRepository) Create(ctx context.Context, user *User) error {
	r.db.Create(&user)
	return nil
}

func (r *DBRepository) Update(ctx context.Context, userID uuid.UUID, user *User) error {
	r.db.Model(&User{}).Where("id = ?", userID).Updates(user)
	return nil
}

func (r *DBRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	r.db.Delete(&User{}, userID)
	return nil
}
