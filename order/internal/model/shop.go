package model

import (
	"github.com/fenky-ng/edot-test-case/order/internal/constant"
	"github.com/google/uuid"
)

type ExtShop struct {
	HttpCode int
	Id       uuid.UUID
	Name     string
	Status   constant.ShopStatus
}

type GetMyShopInput struct {
	JWT string
}

type GetMyShopOutput struct {
	ExtShop
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
