package product

import (
	"database/sql"

	"github.com/leporo/sqlf"
)

type RepoDbProduct struct {
	db  *sql.DB
	sql *sqlf.Dialect
}

type InitRepoDbProductOptions struct {
	DB *sql.DB
}

func InitRepoDbProduct(opts InitRepoDbProductOptions) *RepoDbProduct {
	return &RepoDbProduct{
		db:  opts.DB,
		sql: sqlf.PostgreSQL,
	}
}
