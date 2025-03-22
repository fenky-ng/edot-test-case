package warehouse

import (
	"context"
	"time"

	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
)

func (u *WarehouseUsecase) UpdateWarehouse(ctx context.Context, input model.UpdateWarehouseInput) (output model.UpdateWarehouseOutput, err error) {
	shopOut, err := u.repoHttpShop.GetMyShop(ctx, model.GetMyShopInput{
		JWT: input.JWT,
	})
	if err != nil {
		return output, err
	}

	warehousesOut, err := u.repoDbWarehouse.GetWarehouses(ctx, model.GetWarehousesInput{
		Id: input.WarehouseId,
	})
	if err != nil {
		return output, err
	}
	if len(warehousesOut.Warehouses) == 0 {
		err = in_err.ErrWarehouseNotFound
		return output, err
	}
	if warehousesOut.Warehouses[0].ShopId != shopOut.Id {
		err = in_err.ErrNotWarehouseOwner
		return output, err
	}

	input.UpdatedAt = time.Now().UnixMilli()
	updateOut, err := u.repoDbWarehouse.UpdateWarehouse(ctx, input)
	if err != nil {
		return output, err
	}

	output.Id = updateOut.Id

	return output, nil
}
