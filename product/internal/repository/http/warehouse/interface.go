package warehouse

import (
	"context"

	"github.com/fenky-ng/edot-test-case/product/internal/model"
)

type RepoHttpWarehouseInterface interface {
	GetStocks(ctx context.Context, input model.GetStocksInput) (output model.GetStocksOutput, err error)
}
