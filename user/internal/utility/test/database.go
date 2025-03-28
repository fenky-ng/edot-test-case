package test

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

const (
	LocalDatabaseDSN = "postgres://user:password@localhost:5532/edot_user_db?sslmode=disable"
)

func NewLocalDb() *sql.DB {
	databaseDSN := os.Getenv("DATABASE_URL")
	if databaseDSN == "" {
		databaseDSN = LocalDatabaseDSN
	}

	db, err := sql.Open("postgres", databaseDSN)
	if err != nil {
		panic(err)
	}

	return db
}
