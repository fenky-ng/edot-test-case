package warehouse

import (
	"context"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	dbtx "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/database/tx"
)

type RepoDbWarehouseInterface interface {
	dbtx.DbTxInterface
	InsertWarehouse(ctx context.Context, input model.InsertWarehouseInput) (output model.InsertWarehouseOutput, err error)
	UpdateWarehouse(ctx context.Context, input model.UpdateWarehouseInput) (output model.UpdateWarehouseOutput, err error)
	GetWarehouses(ctx context.Context, input model.GetWarehousesInput) (output model.GetWarehousesOutput, err error)
	GetStocks(ctx context.Context, input model.GetStocksInput) (output model.GetStocksOutput, err error)
	UpsertStock(ctx context.Context, input model.UpsertStockInput) (output model.UpsertStockOutput, err error)
	DeductStock(ctx context.Context, input model.DeductStockInput) (output model.DeductStockOutput, err error)
}
