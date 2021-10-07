package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(5 * time.Microsecond)
	db.SetConnMaxLifetime(9 * time.Second)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(500)
	return db, nil
}
