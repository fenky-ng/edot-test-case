package model

import (
	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	"github.com/google/uuid"
)

type GetProductByIdInput struct {
	Id uuid.UUID
}

type ExtProduct struct {
	Id          uuid.UUID
	Name        string
	Description string
	Price       int64
	Status      constant.ProductStatus
	Shop        ExtShop
}

type GetProductByIdOutput struct {
	ExtProduct
}

type HttpProduct struct {
	Id          uuid.UUID              `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Price       int64                  `json:"price"`
	Status      constant.ProductStatus `json:"status"`
	Shop        HttpShop               `json:"shop"`
}

type HttpGetProductByIdResponse struct {
	Error string `json:"error"`
	HttpProduct
}
