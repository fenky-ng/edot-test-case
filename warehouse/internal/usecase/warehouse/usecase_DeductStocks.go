package warehouse

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	string_util "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/string"
	"github.com/google/uuid"
)

func (u *WarehouseUsecase) DeductStocks(ctx context.Context, input model.DeductStocksInput) (output model.DeductStocksOutput, err error) {
	err = u.warehouseUsecase.ValidateDeductStocks(ctx, input)
	if err != nil {
		return output, err
	}

	ctx, err = u.repoDbWarehouse.Begin(ctx, nil)
	if err != nil {
		return output, err
	}

	defer func() {
		err = u.repoDbWarehouse.CommitOrRollback(ctx, err)
	}()

	for _, stock := range input.Items {
		_, err = u.repoDbWarehouse.DeductStock(ctx, model.DeductStockInput{
			WarehouseId: stock.WarehouseId,
			ProductId:   stock.ProductId,
			Quantity:    stock.Quantity,
			RequestedAt: time.Now().UnixMilli(),
			RequestedBy: input.UserId.String(),
			NoWait:      true,
		})
		if err != nil {
			return output, err
		}
	}

	output.Successful = true

	return output, nil
}

func (u *WarehouseUsecase) ValidateDeductStocks(
	ctx context.Context,
	input model.DeductStocksInput,
) (err error) {
	var warehouseIds []uuid.UUID
	uniqueWarehouseId := make(map[uuid.UUID]struct{})
	for _, item := range input.Items {
		if _, exists := uniqueWarehouseId[item.WarehouseId]; exists {
			continue
		}
		warehouseIds = append(warehouseIds, item.WarehouseId)
		uniqueWarehouseId[item.WarehouseId] = struct{}{}
	}
	warehousesOut, err := u.repoDbWarehouse.GetWarehouses(ctx, model.GetWarehousesInput{
		Ids: warehouseIds,
	})
	if err != nil {
		return err
	}

	var notFoundWarehouseIds []uuid.UUID
	var inactiveWarehouseIds []uuid.UUID
	for _, warehouseId := range warehouseIds {
		if warehouse, exists := warehousesOut.WarehouseById[warehouseId]; !exists {
			notFoundWarehouseIds = append(notFoundWarehouseIds, warehouseId)
		} else if warehouse.Status == constant.WarehouseStatus_Inactive {
			inactiveWarehouseIds = append(inactiveWarehouseIds, warehouseId)
		}
	}

	if len(notFoundWarehouseIds) != 0 {
		err = errors.Join(in_err.ErrWarehouseNotFound, errors.New(
			fmt.Sprintf("warehouse ids: %s", strings.Join(string_util.ParseUuidArrToStringArr(notFoundWarehouseIds), ", ")),
		))
		return err
	}
	if len(inactiveWarehouseIds) != 0 {
		err = errors.Join(in_err.ErrWarehouseInactive, errors.New(
			fmt.Sprintf("warehouse ids: %s", strings.Join(string_util.ParseUuidArrToStringArr(inactiveWarehouseIds), ", ")),
		))
		return err
	}

	return nil
}
