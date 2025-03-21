package database

import (
	"database/sql"
	"log"

	"github.com/fenky-ng/edot-test-case/shop/internal/config"
	_ "github.com/lib/pq"
)

func InitDatabase(cfg *config.Configuration) (db *sql.DB, err error) {
	db, err = connectDatabase("postgres", cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func connectDatabase(dbType string, dbConStr string) (*sql.DB, error) {
	db, err := sql.Open(dbType, dbConStr)
	if err != nil {
		log.Printf("[DB] open database error: %+v", err)
		return nil, err
	}

	log.Println("[DB] database connected")
	return db, nil
}
