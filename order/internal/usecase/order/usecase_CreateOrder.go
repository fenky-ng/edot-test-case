package order

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/fenky-ng/edot-test-case/order/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/order/internal/error"
	"github.com/fenky-ng/edot-test-case/order/internal/model"
	pointer_util "github.com/fenky-ng/edot-test-case/order/internal/utility/pointer"
	"github.com/google/uuid"
)

func (u *OrderUsecase) CreateOrder(ctx context.Context, input model.CreateOrderInput) (output model.CreateOrderOutput, err error) {
	// validate request
	err = u.OrderUsecase.ValidateCreateOrder(ctx, input)
	if err != nil {
		return output, err
	}

	ctx, err = u.repoDbOrder.Begin(ctx, nil)
	if err != nil {
		return output, err
	}

	defer func() {
		err = u.repoDbOrder.CommitOrRollback(ctx, err)
	}()

	// save order and order detail
	now := time.Now()
	orderId := uuid.New()
	orderNo := getOrderNo(now, orderId)
	_, err = u.repoDbOrder.InsertOrder(ctx, model.InsertOrderInput{
		Id:        orderId,
		UserId:    input.UserId,
		OrderNo:   orderNo,
		Status:    constant.OrderStatus_WaitingForPayment,
		CreatedAt: now.UnixMilli(),
		CreatedBy: input.UserId.String(),
	})
	if err != nil {
		return output, err
	}

	output.OrderNo = orderNo
	output.Status = constant.OrderStatus_WaitingForPayment

	_, err = u.repoDbOrder.InsertOrderDetails(ctx, model.InsertOrderDetailsInput{
		OrderId:   orderId,
		Items:     input.Items,
		CreatedAt: now.UnixMilli(),
		CreatedBy: input.UserId.String(),
	})
	if err != nil {
		return output, err
	}

	// deduct stock
	var deductStockItems []model.ExtDeductStockItem
	for _, item := range input.Items {
		deductStockItems = append(deductStockItems, model.ExtDeductStockItem{
			ProductId:   item.ProductId,
			WarehouseId: item.WarehouseId,
			Quantity:    item.Quantity,
		})
	}
	_, err = u.repoHttpWarehouse.DeductStocks(ctx, model.DeductStocksInput{
		UserId:  input.UserId,
		OrderNo: orderNo,
		Items:   deductStockItems,
	})
	if err != nil {
		_, err = u.repoDbOrder.UpdateOrder(ctx, model.UpdateOrderInput{
			OrderNo:      orderNo,
			Status:       pointer_util.PointerOf(constant.OrderStatus_CancelledStockIssue),
			ErrorMessage: pointer_util.PointerOf(err.Error()),
			UpdatedAt:    time.Now().UnixMilli(),
			UpdatedBy:    input.UserId.String(),
		})
		if err == nil {
			output.Status = constant.OrderStatus_CancelledStockIssue
		}

		return output, err
	}

	return output, nil
}

func (u *OrderUsecase) ValidateCreateOrder(
	ctx context.Context,
	input model.CreateOrderInput,
) (err error) {
	shopOut, err := u.repoHttpShop.GetMyShop(ctx, model.GetMyShopInput{
		JWT: input.JWT,
	})
	if err != nil && shopOut.HttpCode != http.StatusNotFound {
		return err
	}

	var productIds []uuid.UUID
	uniqueProductId := make(map[uuid.UUID]struct{})
	for _, item := range input.Items {
		if _, exists := uniqueProductId[item.ProductId]; exists {
			continue
		}
		productIds = append(productIds, item.ProductId)
		uniqueProductId[item.ProductId] = struct{}{}
	}

	productsOut, err := u.repoHttpProduct.GetProducts(ctx, model.GetProductsInput{
		Ids: productIds,
	})
	if err != nil {
		return err
	}

	productById := make(map[uuid.UUID]model.ExtProduct)
	for _, product := range productsOut.Products {
		productById[product.Id] = product
	}

	for index, item := range input.Items {
		product, exists := productById[item.ProductId]
		if !exists {
			err = errors.Join(in_err.ErrProductNotFound, fmt.Errorf("product id: %v", item.ProductId))
			return err
		}

		if product.Status != constant.ProductStatus_Active {
			err = errors.Join(in_err.ErrProductNotActive, fmt.Errorf("product id: %v", item.ProductId))
			return err
		}

		if product.Shop.Id == shopOut.Id {
			err = errors.Join(in_err.ErrUserOwnProduct, fmt.Errorf("product id: %v", item.ProductId))
			return err
		}

		if product.Shop.Status != constant.ShopStatus_Active {
			err = errors.Join(in_err.ErrShopNotActive, fmt.Errorf("product id: %v", item.ProductId))
			return err
		}

		warehouseFound := false
		for _, warehouse := range product.Stock.Warehouses {
			if warehouse.WarehouseId != item.WarehouseId {
				continue
			}

			warehouseFound = true

			if warehouse.WarehouseStatus != constant.WarehouseStatus_Active {
				err = errors.Join(in_err.ErrWarehouseNotActive, fmt.Errorf("product id: %v, warehouse id: %v", item.ProductId, item.WarehouseId))
				return err
			}

			if warehouse.Stock < item.Quantity {
				err = errors.Join(in_err.ErrInsufficientStock, fmt.Errorf("product id: %v, warehouse id: %v", item.ProductId, item.WarehouseId))
				return err
			}
		}

		if !warehouseFound {
			err = errors.Join(in_err.ErrInsufficientStock, fmt.Errorf("product id: %v, warehouse id: %v", item.ProductId, item.WarehouseId))
			return err
		}

		if item.Quantity < 1 {
			err = errors.Join(in_err.ErrInvalidOrderQuantity, fmt.Errorf("product id: %v, warehouse id: %v", item.ProductId, item.WarehouseId))
			return err
		}

		input.Items[index].Price = product.Price
	}

	return nil
}

func getOrderNo(
	timestamp time.Time,
	orderId uuid.UUID,
) string {
	return fmt.Sprintf("ORD/%s/%s", timestamp.Format("20060102150405"), orderId.String())
}
