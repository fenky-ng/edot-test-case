package model

import (
	"github.com/google/uuid"

	"github.com/fenky-ng/edot-test-case/order/internal/constant"
)

type ExtProduct struct {
	Id          uuid.UUID
	Name        string
	Description string
	Price       int64
	Status      constant.ProductStatus
	Shop        ExtShop
	Stock       ExtProductStock
}

type ExtProductStock struct {
	Total      int64
	Warehouses []ExtProductWarehouse
}

type ExtProductWarehouse struct {
	WarehouseId     uuid.UUID
	WarehouseStatus constant.WarehouseStatus
	Stock           int64
}

type GetProductsInput struct {
	Ids []uuid.UUID
}

type GetProductsOutput struct {
	Products []ExtProduct
}

type HttpProduct struct {
	Id          uuid.UUID              `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Price       int64                  `json:"price"`
	Status      constant.ProductStatus `json:"status"`
	Shop        HttpShop               `json:"shop"`
	Stock       HttpProductStock       `json:"stock"`
}

type HttpProductStock struct {
	Total      int64                  `json:"total"`
	Warehouses []HttpProductWarehouse `json:"warehouses"`
}

type HttpProductWarehouse struct {
	WarehouseId     uuid.UUID                `json:"warehouseId"`
	WarehouseStatus constant.WarehouseStatus `json:"warehouseStatus"`
	Stock           int64                    `json:"stock"`
}

type HttpGetProductsResponse struct {
	Error string        `json:"error"`
	Data  []HttpProduct `json:"data"`
}
