package model

import (
	"github.com/fenky-ng/edot-test-case/product/internal/constant"
	"github.com/google/uuid"
)

type ExtShop struct {
	Id     uuid.UUID
	Name   string
	Status constant.ShopStatus
}

type GetMyShopInput struct {
	JWT string
}

type GetMyShopOutput struct {
	ExtShop
}

type GetShopsInput struct {
	Ids []uuid.UUID
}

type GetShopsOutput struct {
	Shops []ExtShop
}

type HttpShop struct {
	Id     uuid.UUID           `json:"id"`
	Name   string              `json:"name"`
	Status constant.ShopStatus `json:"status"`
}

type HttpGetMyShopResponse struct {
	Error string `json:"error"`
	HttpShop
}

type HttpGetShopsResponse struct {
	Error string     `json:"error"`
	Data  []HttpShop `json:"data"`
}
