package model

import (
	"github.com/fenky-ng/edot-test-case/shop/internal/constant"
	"github.com/google/uuid"
)

type RestAPIErrorResponse struct {
	Error string `json:"error"`
}

type RestAPICreateShopRequest struct {
	Name string `json:"name"`
}

type RestAPICreateShopResponse struct {
	Id uuid.UUID `json:"id"`
}

type RestAPIShop struct {
	Id      uuid.UUID           `json:"id"`
	OwnerId uuid.UUID           `json:"ownerId"`
	Name    string              `json:"name"`
	Status  constant.ShopStatus `json:"status"`
}

type RestAPIGetMyShopResponse struct {
	RestAPIShop
}

type RestAPIGetShopByIdResponse struct {
	RestAPIShop
}

type RestAPIGetShopsResponse struct {
	Shops []RestAPIShop `json:"data"`
}
