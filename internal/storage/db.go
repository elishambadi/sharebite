package storage

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func Connect() (*sql.DB, error) {
	connStr := "user=postgres dbname=sharebite sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
