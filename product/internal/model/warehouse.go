package model

import (
	"github.com/fenky-ng/edot-test-case/product/internal/constant"
	"github.com/google/uuid"
)

type ExtProductStock struct {
	ProductId  uuid.UUID
	Warehouses []ExtProductWarehouse
}

type ExtProductWarehouse struct {
	WarehouseId     uuid.UUID
	WarehouseStatus constant.WarehouseStatus
	Stock           int64
}

type GetStocksInput struct {
	ProductIds []uuid.UUID
}

type GetStocksOutput struct {
	Stocks []ExtProductStock
}

type HttpProductStock struct {
	ProductId  uuid.UUID              `json:"productId"`
	Warehouses []HttpProductWarehouse `json:"warehouses"`
}

type HttpProductWarehouse struct {
	WarehouseId     uuid.UUID                `json:"warehouseId"`
	WarehouseStatus constant.WarehouseStatus `json:"warehouseStatus"`
	Stock           int64                    `json:"stock"`
}

type HttpGetStocksResponse struct {
	Error string             `json:"error"`
	Data  []HttpProductStock `json:"data"`
}
