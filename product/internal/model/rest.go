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
	Stock       RestAPIStock           `json:"stock"`
}

type RestAPIShop struct {
	Id     uuid.UUID           `json:"id"`
	Name   string              `json:"name"`
	Status constant.ShopStatus `json:"status"`
}

type RestAPIStock struct {
	Total      int                       `json:"total"`
	Warehouses []RestAPIProductWarehouse `json:"warehouses"`
}

type RestAPIProductWarehouse struct {
	Id     uuid.UUID                    `json:"id"`
	Stock  int                          `json:"stock"`
	Status constant.ShopWarehouseStatus `json:"statu"`
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
