package warehouse

import (
	"context"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
)

func (u *WarehouseUsecase) GetMyWarehouses(ctx context.Context, input model.GetMyWarehousesInput) (output model.GetMyWarehousesOutput, err error) {
	shopOut, err := u.repoHttpShop.GetMyShop(ctx, model.GetMyShopInput{
		JWT: input.JWT,
	})
	if err != nil {
		return output, err
	}

	warehousesOut, err := u.repoDbWarehouse.GetWarehouses(ctx, model.GetWarehousesInput{
		ShopId: shopOut.Id,
	})
	if err != nil {
		return output, err
	}

	output.Warehouses = warehousesOut.Warehouses

	return output, nil
}
