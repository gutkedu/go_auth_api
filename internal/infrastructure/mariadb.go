package infrastructure

import (
	"database/sql"
	"time"
)

func ConnectToMariaDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(mariadb:3306)/fiber_golangapi")
	if err != nil {
		return nil, err
	}
	// Set up important parts as was told by the documentation.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
