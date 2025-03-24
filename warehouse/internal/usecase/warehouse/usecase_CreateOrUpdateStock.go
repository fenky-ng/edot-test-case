package warehouse

import (
	"context"
	"errors"
	"time"

	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	"github.com/google/uuid"
)

func (u *WarehouseUsecase) CreateOrUpdateStock(ctx context.Context, input model.CreateOrUpdateStockInput) (output model.CreateOrUpdateStockOutput, err error) {
	err = u.WarehouseUsecase.ValidateCreateOrUpdateStock(ctx, input)
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

	if input.ToWarehouseId == uuid.Nil { // SET STOCK
		_, err = u.repoDbWarehouse.UpsertStock(ctx, model.UpsertStockInput{
			Id:          uuid.New(),
			WarehouseId: input.WarehouseId,
			ProductId:   input.ProductId,
			Stock:       input.Stock,
			IsTransfer:  false,
			UpsertedAt:  time.Now().UnixMilli(),
			UpsertedBy:  input.UserId.String(),
		})
		if err != nil {
			return output, err
		}

		output.Successful = true

		return output, nil
	}

	// TRANSFER
	err = u.WarehouseUsecase.TransferStock(ctx, input)
	if err != nil {
		return output, err
	}

	output.Successful = true

	return output, nil
}

func (u *WarehouseUsecase) ValidateCreateOrUpdateStock(
	ctx context.Context,
	input model.CreateOrUpdateStockInput,
) (err error) {
	shopOut, err := u.repoHttpShop.GetMyShop(ctx, model.GetMyShopInput{
		JWT: input.JWT,
	})
	if err != nil {
		return err
	}

	// validate warehouse
	isTransfer := input.ToWarehouseId != uuid.Nil

	warehouseIds := []uuid.UUID{
		input.WarehouseId,
	}
	if isTransfer {
		warehouseIds = append(warehouseIds, input.ToWarehouseId)
	}
	warehousesOut, err := u.repoDbWarehouse.GetWarehouses(ctx, model.GetWarehousesInput{
		Ids: warehouseIds,
	})
	if err != nil {
		return err
	}

	err = validateWarehouseExistenceAndOwnership(shopOut.Id, input.WarehouseId, warehousesOut.WarehouseById)
	if err != nil {
		return err
	}

	if isTransfer {
		err = validateWarehouseExistenceAndOwnership(shopOut.Id, input.ToWarehouseId, warehousesOut.WarehouseById)
		if err != nil {
			return err
		}
	}

	// validate product
	productOut, err := u.repoHttpProduct.GetProductById(ctx, model.GetProductByIdInput{
		Id: input.ProductId,
	})
	if err != nil {
		return err
	}

	if productOut.Shop.Id != shopOut.Id {
		err = in_err.ErrNotProductOwner
		return err
	}

	return nil
}

func validateWarehouseExistenceAndOwnership(
	shopId uuid.UUID,
	warehouseId uuid.UUID,
	warehouseById map[uuid.UUID]model.Warehouse,
) (err error) {
	if warehouse, exists := warehouseById[warehouseId]; !exists {
		err = errors.Join(in_err.ErrWarehouseNotFound, errors.New("warehouse id: "+warehouseId.String()))
		return err
	} else if warehouse.ShopId != shopId {
		err = errors.Join(in_err.ErrNotWarehouseOwner, errors.New("warehouse id: "+warehouseId.String()))
		return err
	}
	return nil
}

func (u *WarehouseUsecase) TransferStock(
	ctx context.Context,
	input model.CreateOrUpdateStockInput,
) (err error) {
	if input.WarehouseId == input.ToWarehouseId {
		err = in_err.ErrInvalidStockTransferDestination
		return err
	}

	now := time.Now()

	_, err = u.repoDbWarehouse.DeductStock(ctx, model.DeductStockInput{
		WarehouseId: input.WarehouseId,
		ProductId:   input.ProductId,
		Quantity:    input.Stock,
		RequestedAt: now.UnixMilli(),
		RequestedBy: input.UserId.String(),
	})
	if err != nil {
		return err
	}

	_, err = u.repoDbWarehouse.UpsertStock(ctx, model.UpsertStockInput{
		Id:          uuid.New(),
		WarehouseId: input.ToWarehouseId,
		ProductId:   input.ProductId,
		Stock:       input.Stock,
		IsTransfer:  true,
		UpsertedAt:  now.UnixMilli(),
		UpsertedBy:  input.UserId.String(),
	})
	if err != nil {
		return err
	}

	return nil
}
