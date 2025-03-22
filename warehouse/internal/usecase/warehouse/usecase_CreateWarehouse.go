package warehouse

import (
	"context"
	"time"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	"github.com/google/uuid"
)

func (u *WarehouseUsecase) CreateWarehouse(ctx context.Context, input model.CreateWarehouseInput) (output model.CreateWarehouseOutput, err error) {
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

	if len(warehousesOut.Warehouses) >= constant.MaxWarehousePerShop {
		err = in_err.ErrMaxWarehousePerShop
		return output, err
	}

	insertOut, err := u.repoDbWarehouse.InsertWarehouse(ctx, model.InsertWarehouseInput{
		Id:        uuid.New(),
		ShopId:    shopOut.Id,
		Name:      input.Name,
		Status:    input.Status,
		CreatedAt: time.Now().UnixMilli(),
		CreatedBy: input.UserId.String(),
	})
	if err != nil {
		return output, err
	}

	output.Id = insertOut.Id

	return output, nil
}
