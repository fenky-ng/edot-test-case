package warehouse

import (
	"database/sql"

	dbtx "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/database/tx"
	"github.com/leporo/sqlf"
)

type RepoDbWarehouse struct {
	dbtx.DbTxInterface
	db  *sql.DB
	sql *sqlf.Dialect
}

type InitRepoDbWarehouseOptions struct {
	DbTx dbtx.DbTxInterface
	DB   *sql.DB
}

func InitRepoDbWarehouse(opts InitRepoDbWarehouseOptions) *RepoDbWarehouse {
	return &RepoDbWarehouse{
		DbTxInterface: opts.DbTx,
		db:            opts.DB,
		sql:           sqlf.PostgreSQL,
	}
}
