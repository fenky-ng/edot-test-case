package model

import (
	"github.com/fenky-ng/edot-test-case/product/internal/constant"
	"github.com/google/uuid"
)

type RestAPIErrorResponse struct {
	Error string `json:"error"`
}

type RestAPICreateProductRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Price       int64                  `json:"price"`
	Status      constant.ProductStatus `json:"status"`
}

type RestAPICreateProductResponse struct {
	Id uuid.UUID `json:"id"`
}

type RestAPIProduct struct {
	Id          uuid.UUID              `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Price       int64                  `json:"price"`
	Status      constant.ProductStatus `json:"status"`
	Shop        RestAPIShop            `json:"shop"`
	Stock       RestAPIProductStock    `json:"stock"`
}

type RestAPIShop struct {
	Id     uuid.UUID           `json:"id"`
	Name   string              `json:"name"`
	Status constant.ShopStatus `json:"status"`
}

type RestAPIProductStock struct {
	Total      int64                     `json:"total"`
	Warehouses []RestAPIProductWarehouse `json:"warehouses"`
}

type RestAPIProductWarehouse struct {
	WarehouseId     uuid.UUID                `json:"warehouseId"`
	WarehouseStatus constant.WarehouseStatus `json:"warehouseStatus"`
	Stock           int64                    `json:"stock"`
}

type RestAPIGetMyProductsResponse struct {
	Products []RestAPIProduct `json:"data"`
}

type RestAPIGetProductsResponse struct {
	Products []RestAPIProduct `json:"data"`
}

type RestAPIGetProductByIdResponse struct {
	RestAPIProduct
}
