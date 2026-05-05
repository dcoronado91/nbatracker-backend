package db

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
    connStr := "postgres://user:password@localhost:5432/nbadb?sslmode=disable"
    return sql.Open("postgres", connStr)
}