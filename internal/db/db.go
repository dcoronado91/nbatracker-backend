package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	return sql.Open("postgres", connStr)
}
