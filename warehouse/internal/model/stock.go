package model

import (
	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	"github.com/google/uuid"
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

type UpsertStockOutput struct {
	Id uuid.UUID
}

type DeductStockInput struct {
	WarehouseId uuid.UUID
	ProductId   uuid.UUID
	Quantity    int64
	RequestedAt int64
	RequestedBy string
	NoWait      bool
}

type DeductStockOutput struct {
	Successful bool
}
