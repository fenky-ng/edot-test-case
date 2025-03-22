package warehouse

import (
	"context"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
)

type RepoDbWarehouseInterface interface {
	InsertWarehouse(ctx context.Context, input model.InsertWarehouseInput) (output model.InsertWarehouseOutput, err error)
	UpdateWarehouse(ctx context.Context, input model.UpdateWarehouseInput) (output model.UpdateWarehouseOutput, err error)
	GetWarehouses(ctx context.Context, input model.GetWarehousesInput) (output model.GetWarehousesOutput, err error)
	GetStocks(ctx context.Context, input model.GetStocksInput) (output model.GetStocksOutput, err error)
}
