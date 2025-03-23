package tx

import (
	"context"
	"database/sql"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
)

func (r *DbTx) Begin(ctx context.Context, opts *sql.TxOptions) (context.Context, error) {
	tx, ok := r.txFromContext(ctx)
	if ok {
		return ctx, nil
	}

	tx, err := r.db.BeginTx(ctx, opts)
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, constant.DbTxTransactionKey, tx), nil
}

func (r *DbTx) txFromContext(ctx context.Context) (tx *sql.Tx, ok bool) {
	tx, ok = ctx.Value(constant.DbTxTransactionKey).(*sql.Tx)
	return tx, ok
}
