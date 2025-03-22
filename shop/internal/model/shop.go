package model

import (
	"github.com/fenky-ng/edot-test-case/shop/internal/constant"
	"github.com/google/uuid"
)

type Shop struct {
	Id      uuid.UUID
	OwnerId uuid.UUID
	Name    string
	Status  constant.ShopStatus
}

type CreateShopInput struct {
	UserId uuid.UUID
	Name   string
}

type CreateShopOutput struct {
	Id uuid.UUID
}

type GetMyShopInput struct {
	UserId uuid.UUID
}

type GetMyShopOutput struct {
	Shop
}

type GetShopsInput struct {
	Ids []uuid.UUID
}

type GetShopsOutput struct {
	Shops []Shop
}

type GetShopByIdInput struct {
	Id uuid.UUID
}

type GetShopByIdOutput struct {
	Shop
}

type GetShopInput struct {
	ShopId         uuid.UUID
	OwnerId        uuid.UUID
	ExcludeDeleted bool
}

type GetShopOutput struct {
	Id        uuid.UUID
	OwnerId   uuid.UUID
	Name      string
	Status    constant.ShopStatus
	DeletedAt int64
}

type InsertShopInput struct {
	Id        uuid.UUID
	OwnerId   uuid.UUID
	Name      string
	Status    constant.ShopStatus
	CreatedAt int64
	CreatedBy string
}

type InsertShopOutput struct {
	Id uuid.UUID
}
