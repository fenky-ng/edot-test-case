package model

import (
	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	"github.com/google/uuid"
)

type Warehouse struct {
	Id     uuid.UUID
	ShopId uuid.UUID
	Name   string
	Status constant.WarehouseStatus
}

type CreateWarehouseInput struct {
	JWT    string
	UserId uuid.UUID
	Name   string
	Status constant.WarehouseStatus
}

type CreateWarehouseOutput struct {
	Id uuid.UUID
}

type UpdateWarehouseInput struct {
	JWT         string
	UserId      uuid.UUID
	WarehouseId uuid.UUID
	Name        *string
	Status      *constant.WarehouseStatus
	UpdatedAt   int64
}

type UpdateWarehouseOutput struct {
	Id uuid.UUID
}

type GetMyWarehousesInput struct {
	JWT    string
	UserId uuid.UUID
}

type GetMyWarehousesOutput struct {
	Warehouses []Warehouse
}

type InsertWarehouseInput struct {
	Id        uuid.UUID
	ShopId    uuid.UUID
	Name      string
	Status    constant.WarehouseStatus
	CreatedAt int64
	CreatedBy string
}

type InsertWarehouseOutput struct {
	Id uuid.UUID
}

type GetWarehousesInput struct {
	Ids    []uuid.UUID
	ShopId uuid.UUID
}

type GetWarehousesOutput struct {
	Warehouses    []Warehouse
	WarehouseById map[uuid.UUID]Warehouse
}
