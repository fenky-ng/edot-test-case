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

type DeductStockInput struct {
	// TODO
}

type DeductStockOutput struct {
	// TODO
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
