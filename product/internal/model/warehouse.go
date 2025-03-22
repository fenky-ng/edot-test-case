package model

import (
	"github.com/fenky-ng/edot-test-case/product/internal/constant"
	"github.com/google/uuid"
)

type ExtProductWarehouse struct {
	ProductId   uuid.UUID
	WarehouseId uuid.UUID
	Stock       int
	Status      constant.ShopWarehouseStatus
}

type GetProductWarehousesInput struct {
	ProductIds []uuid.UUID
}

type GetProductWarehousesOutput struct {
	WarehousesByProductId map[uuid.UUID][]ExtProductWarehouse
}

type HttpWarehouse struct {
	ProductId   uuid.UUID
	WarehouseId uuid.UUID
	Stock       int
	Status      constant.ShopWarehouseStatus
}

type HttpGetProductWarehousesResponse struct {
	Error string                                             `json:"error"`
	Data  map[uuid.UUID]HttpGetProductWarehousesResponseItem `json:"data"`
}

type HttpGetProductWarehousesResponseItem struct {
	Warehouses []HttpWarehouse `json:"warehouses"`
}
