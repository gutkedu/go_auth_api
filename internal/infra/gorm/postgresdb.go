package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPgDB() (*gorm.DB, error) {
	dsn := "postgres://golangapi:golang@localhost:5432/golang_database"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
