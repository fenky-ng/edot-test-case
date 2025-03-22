package model

import (
	"github.com/fenky-ng/edot-test-case/product/internal/constant"
	"github.com/google/uuid"
)

type Shop struct {
	Id     uuid.UUID
	Name   string
	Status constant.ShopStatus
}

type Stock struct {
	Total      int64
	Warehouses []ProductWarehouse
}

type ProductWarehouse struct {
	WarehouseId     uuid.UUID
	WarehouseStatus constant.WarehouseStatus
	Stock           int64
}

type Product struct {
	Id          uuid.UUID
	Name        string
	Description string
	Status      constant.ProductStatus
	Price       int64
	Shop        Shop
	Stock       Stock
}

type CreateProductInput struct {
	JWT         string
	UserId      uuid.UUID
	Name        string
	Description string
	Price       int64
	Status      constant.ProductStatus
}

type CreateProductOutput struct {
	Id uuid.UUID
}

type GetMyProductsInput struct {
	JWT    string
	UserId uuid.UUID
}

type GetMyProductsOutput struct {
	Products []Product
}

type GetProductsInput struct {
	ShopId uuid.UUID
}

type GetProductsOutput struct {
	Products []Product
}

type GetProductByIdInput struct {
	Id uuid.UUID
}

type GetProductByIdOutput struct {
	Product Product
}

type InsertProductInput struct {
	Id          uuid.UUID
	ShopId      uuid.UUID
	Name        string
	Description string
	Price       int64
	Status      constant.ProductStatus
	CreatedAt   int64
	CreatedBy   string
}

type InsertProductOutput struct {
	Id uuid.UUID
}

type GetProductInput struct {
	Id uuid.UUID
}

type GetProductOutput struct {
	Product Product
}
