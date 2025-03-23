package tx

import (
	"database/sql"
)

type DbTx struct {
	db *sql.DB
}

type InitDbTxOptions struct {
	DB *sql.DB
}

func InitDbTx(opts InitDbTxOptions) *DbTx {
	return &DbTx{
		db: opts.DB,
	}
}
