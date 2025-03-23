package tx

import (
	"context"
	"database/sql"

	"github.com/leporo/sqlf"
)

type DbTxInterface interface {
	Begin(ctx context.Context, opts *sql.TxOptions) (context.Context, error)
	CommitOrRollback(ctx context.Context, err error) error
	UseTx(ctx context.Context) sqlf.Executor
}
