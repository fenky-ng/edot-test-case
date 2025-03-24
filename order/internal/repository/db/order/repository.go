package order

import (
	"database/sql"

	dbtx "github.com/fenky-ng/edot-test-case/order/internal/utility/database/tx"
	"github.com/leporo/sqlf"
)

type RepoDbOrder struct {
	dbtx.DbTxInterface
	db  *sql.DB
	sql *sqlf.Dialect
}

type InitRepoDbOrderOptions struct {
	DbTx dbtx.DbTxInterface
	DB   *sql.DB
}

func InitRepoDbOrder(opts InitRepoDbOrderOptions) *RepoDbOrder {
	return &RepoDbOrder{
		DbTxInterface: opts.DbTx,
		db:            opts.DB,
		sql:           sqlf.PostgreSQL,
	}
}
