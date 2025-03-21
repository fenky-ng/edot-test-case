package shop

import (
	"database/sql"

	"github.com/leporo/sqlf"
)

type RepoDbShop struct {
	db  *sql.DB
	sql *sqlf.Dialect
}

type InitRepoDbShopOptions struct {
	DB *sql.DB
}

func InitRepoDbShop(opts InitRepoDbShopOptions) *RepoDbShop {
	return &RepoDbShop{
		db:  opts.DB,
		sql: sqlf.PostgreSQL,
	}
}
