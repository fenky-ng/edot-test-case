package model

import (
	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	"github.com/google/uuid"
)

type RestAPIErrorResponse struct {
	Error string `json:"error"`
}

type RestAPICreateWarehouseRequest struct {
	Name   string                   `json:"name"`
	Status constant.WarehouseStatus `json:"status"`
}

type RestAPICreateWarehouseResponse struct {
	Id uuid.UUID `json:"id"`
}

type RestAPIUpdateWarehouseRequest struct {
	Name   *string                   `json:"name,omitempty"`
	Status *constant.WarehouseStatus `json:"status,omitempty"`
}

type RestAPIUpdateWarehouseResponse struct {
	Id uuid.UUID `json:"id"`
}

type RestAPIWarehouse struct {
	Id     uuid.UUID                `json:"id"`
	ShopId uuid.UUID                `json:"shopId"`
	Name   string                   `json:"name"`
	Status constant.WarehouseStatus `json:"status"`
}

type RestAPIGetMyWarehousesResponse struct {
	Warehouses []RestAPIWarehouse `json:"data"`
}

type RestAPICreateOrUpdateStockRequest struct {
	WarehouseId   uuid.UUID `json:"warehouseId"`
	ProductId     uuid.UUID `json:"productId"`
	Stock         int64     `json:"stock"`
	ToWarehouseId uuid.UUID `json:"toWarehouseId"` // Optional, only for transfer
}

type RestAPICreateOrUpdateStockResponse struct {
	Successful bool `json:"successful"`
}

type RestAPIProductStock struct {
	ProductId  uuid.UUID                 `json:"productId"`
	Warehouses []RestAPIProductWarehouse `json:"warehouses"`
}

type RestAPIProductWarehouse struct {
	WarehouseId     uuid.UUID                `json:"warehouseId"`
	WarehouseStatus constant.WarehouseStatus `json:"warehouseStatus"`
	Stock           int64                    `json:"stock"`
}

type RestAPIGetStocksResponse struct {
	Products []RestAPIProductStock `json:"data"`
}

type RestAPIDeductStocksRequest struct {
	UserId  uuid.UUID                `json:"userId"`
	OrderNo string                   `json:"orderNo"`
	Items   []RestAPIDeductStockItem `json:"items"`
	Release bool                     `json:"release"`
}

type RestAPIDeductStockItem struct {
	ProductId   uuid.UUID `json:"productId"`
	WarehouseId uuid.UUID `json:"warehouseId"`
	Quantity    int64     `json:"quantity"`
}

type RestAPIDeductStocksResponse struct {
	Successful bool `json:"successful"`
}
