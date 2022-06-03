package migrations

import (
	"github.com/gutkedu/golang_api/internal/modules/roles"
	"github.com/gutkedu/golang_api/internal/modules/user"
	"gorm.io/gorm"
)

func RunGormMigrations(dbConn *gorm.DB) {
	dbConn.AutoMigrate(&user.User{})
	dbConn.AutoMigrate(&roles.Role{})
}
