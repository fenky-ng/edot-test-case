package model

import (
	"fmt"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

type ProductStock struct {
	ProductId  uuid.UUID
	Warehouses []ProductWarehouse
}

type ProductWarehouse struct {
	WarehouseId     uuid.UUID
	WarehouseStatus constant.WarehouseStatus
	Stock           int64
}

type CreateOrUpdateStockInput struct {
	JWT           string
	UserId        uuid.UUID
	WarehouseId   uuid.UUID
	ProductId     uuid.UUID
	Stock         int64
	ToWarehouseId uuid.UUID
}

type CreateOrUpdateStockOutput struct {
	Successful bool
}

type GetProductStocksInput struct {
	ProductIds []uuid.UUID
}

type GetProductStocksOutput struct {
	ProductStocks []ProductStock
}

type DeductStockItem struct {
	ProductId   uuid.UUID
	WarehouseId uuid.UUID
	Quantity    int64
}

type DeductStocksInput struct {
	UserId  uuid.UUID
	OrderNo string
	Items   []DeductStockItem
	Release bool
}

type DeductStocksOutput struct {
	Successful bool
}

type GetStocksInput struct {
	ProductIds []uuid.UUID
}

type Stock struct {
	Id              uuid.UUID
	ProductId       uuid.UUID
	WarehouseId     uuid.UUID
	WarehouseStatus constant.WarehouseStatus
	Stock           int64
}

type GetStocksOutput struct {
	Stocks            []Stock
	StocksByProductId map[uuid.UUID][]Stock
}

type UpsertStockInput struct {
	Id          uuid.UUID
	WarehouseId uuid.UUID
	ProductId   uuid.UUID
	Stock       int64
	IsTransfer  bool
	UpsertedAt  int64
	UpsertedBy  string
}

func (expectedInput UpsertStockInput) Matcher() gomock.Matcher {
	return gomock.Cond(func(x any) bool {
		actualInput := x.(UpsertStockInput)

		// set zero value for ignored attributes
		expectedInput.Id, actualInput.Id = uuid.Nil, uuid.Nil
		expectedInput.UpsertedAt, actualInput.UpsertedAt = 0, 0

		diff := cmp.Diff(expectedInput, actualInput)
		if diff != "" {
			fmt.Printf("[UpsertStockInputMatcher] DEBUG input mismatch (-want +got):\n%s\n", diff)
		}

		return diff == ""
	})
}

type UpsertStockOutput struct {
	Id uuid.UUID
}

type DeductStockInput struct {
	WarehouseId uuid.UUID
	ProductId   uuid.UUID
	Quantity    int64
	Release     bool
	RequestedAt int64
	RequestedBy string
	NoWait      bool
}

func (expectedInput DeductStockInput) Matcher() gomock.Matcher {
	return gomock.Cond(func(x any) bool {
		actualInput := x.(DeductStockInput)

		// set zero value for ignored attributes
		expectedInput.RequestedAt, actualInput.RequestedAt = 0, 0

		diff := cmp.Diff(expectedInput, actualInput)
		if diff != "" {
			fmt.Printf("[DeductStockInputMatcher] DEBUG input mismatch (-want +got):\n%s\n", diff)
		}

		return diff == ""
	})
}

type DeductStockOutput struct {
	Successful bool
}
