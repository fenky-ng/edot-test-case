package warehouse

import (
	"database/sql"

	"github.com/leporo/sqlf"
)

type RepoDbWarehouse struct {
	db  *sql.DB
	sql *sqlf.Dialect
}

type InitRepoDbWarehouseOptions struct {
	DB *sql.DB
}

func InitRepoDbWarehouse(opts InitRepoDbWarehouseOptions) *RepoDbWarehouse {
	return &RepoDbWarehouse{
		db:  opts.DB,
		sql: sqlf.PostgreSQL,
	}
}
