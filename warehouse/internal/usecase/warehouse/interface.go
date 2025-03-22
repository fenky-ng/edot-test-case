package warehouse

import (
	"context"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
)

type WarehouseUsecaseInterface interface {
	CreateWarehouse(ctx context.Context, input model.CreateWarehouseInput) (output model.CreateWarehouseOutput, err error)
	UpdateWarehouse(ctx context.Context, input model.UpdateWarehouseInput) (output model.UpdateWarehouseOutput, err error)
	GetMyWarehouses(ctx context.Context, input model.GetMyWarehousesInput) (output model.GetMyWarehousesOutput, err error)
	CreateOrUpdateStock(ctx context.Context, input model.CreateOrUpdateStockInput) (output model.CreateOrUpdateStockOutput, err error)
	GetProductStocks(ctx context.Context, input model.GetProductStocksInput) (output model.GetProductStocksOutput, err error)
	DeductStock(ctx context.Context, input model.DeductStockInput) (output model.DeductStockOutput, err error)
}
