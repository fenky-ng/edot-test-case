package tx

import (
	"context"

	"github.com/leporo/sqlf"
)

func (r *DbTx) UseTx(ctx context.Context) sqlf.Executor {
	tx, ok := r.txFromContext(ctx)
	if ok {
		return tx
	}
	return r.db
}
