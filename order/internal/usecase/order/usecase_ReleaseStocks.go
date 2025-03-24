package order

import (
	"context"

	"github.com/fenky-ng/edot-test-case/order/internal/model"
)

func (u *OrderUsecase) ReleaseStocks(ctx context.Context, input model.ReleaseStocksInput) (output model.ReleaseStocksOutput, err error) {
	orderDetailsOut, err := u.repoDbOrder.GetOrderDetails(ctx, model.GetOrderDetailsInput{
		OrderNo: input.OrderNo,
	})
	if err != nil {
		return output, err
	}

	var releaseStockItems []model.ExtDeductStockItem
	for _, item := range orderDetailsOut.Items {
		releaseStockItems = append(releaseStockItems, model.ExtDeductStockItem{
			ProductId:   item.ProductId,
			WarehouseId: item.WarehouseId,
			Quantity:    item.Quantity,
		})
	}

	_, err = u.repoHttpWarehouse.DeductStocks(ctx, model.DeductStocksInput{
		UserId:  input.UserId,
		OrderNo: input.OrderNo,
		Items:   releaseStockItems,
		Release: true,
	})
	if err != nil {
		return output, err
	}

	return output, nil
}
