package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMariaDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(mariadb)/golang_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
