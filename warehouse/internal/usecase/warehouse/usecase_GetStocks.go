package warehouse

import (
	"context"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
)

func (u *WarehouseUsecase) GetProductStocks(ctx context.Context, input model.GetProductStocksInput) (output model.GetProductStocksOutput, err error) {
	stocksOut, err := u.repoDbWarehouse.GetStocks(ctx, model.GetStocksInput{
		ProductIds: input.ProductIds,
	})
	if err != nil {
		return output, err
	}

	for productId, stocks := range stocksOut.StocksByProductId {
		productStock := model.ProductStock{
			ProductId:  productId,
			Warehouses: make([]model.ProductWarehouse, 0),
		}

		for _, stock := range stocks {
			productStock.Warehouses = append(productStock.Warehouses, model.ProductWarehouse{
				WarehouseId:     stock.WarehouseId,
				WarehouseStatus: stock.WarehouseStatus,
				Stock:           stock.Stock,
			})
		}

		output.ProductStocks = append(output.ProductStocks, productStock)
	}

	return output, nil
}
