package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMariaDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(mariadb:3306)/golang_api_database"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
