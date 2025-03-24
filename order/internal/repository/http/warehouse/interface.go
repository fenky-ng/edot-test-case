package warehouse

import (
	"context"

	"github.com/fenky-ng/edot-test-case/order/internal/model"
)

type RepoHttpWarehouseInterface interface {
	DeductStocks(ctx context.Context, input model.DeductStocksInput) (output model.DeductStocksOutput, err error)
}
