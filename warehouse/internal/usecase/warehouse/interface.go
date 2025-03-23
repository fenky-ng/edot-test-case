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
	DeductStocks(ctx context.Context, input model.DeductStocksInput) (output model.DeductStocksOutput, err error)

	// INTERNAL - warehouse usecase related function
	ValidateCreateOrUpdateStock(ctx context.Context, input model.CreateOrUpdateStockInput) (err error)
	TransferStock(ctx context.Context, input model.CreateOrUpdateStockInput) (err error)
	ValidateDeductStocks(ctx context.Context, input model.DeductStocksInput) (err error)
}
