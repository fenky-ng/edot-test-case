package model

import (
	"github.com/fenky-ng/edot-test-case/order/internal/constant"
	"github.com/google/uuid"
)

type Order struct {
	Id           uuid.UUID
	UserId       uuid.UUID
	OrderNo      string
	Status       constant.OrderStatus
	PaymentRefNo string
}

type OrderItem struct {
	ProductId   uuid.UUID
	WarehouseId uuid.UUID
	Quantity    int64
	Price       int64
}

type CreateOrderInput struct {
	JWT    string
	UserId uuid.UUID
	Items  []OrderItem
}

type CreateOrderOutput struct {
	OrderNo string
	Status  constant.OrderStatus
}

type InsertOrderInput struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	OrderNo   string
	Status    constant.OrderStatus
	CreatedAt int64
	CreatedBy string
}

type InsertOrderOutput struct {
	Id uuid.UUID
}

type InsertOrderDetailsInput struct {
	OrderId   uuid.UUID
	Items     []OrderItem
	CreatedAt int64
	CreatedBy string
}

type InsertOrderDetailsOutput struct {
	Successful bool
}

type UpdateOrderInput struct {
	OrderNo      string
	Status       *constant.OrderStatus
	PaymentRefNo *string
	ErrorMessage *string
	UpdatedAt    int64
	UpdatedBy    string
}

type UpdateOrderOutput struct {
	Successful bool
}

type GetOrdersInput struct {
	UserId  uuid.UUID
	OrderNo string

	// internal
	GetExpiredOrders bool
}

type GetOrdersOutput struct {
	Orders []Order
}

type ReleaseStocksInput struct {
	UserId  uuid.UUID
	OrderNo string
}

type ReleaseStocksOutput struct {
	Successful bool
}

type GetOrderDetailsInput struct {
	OrderNo string
}

type GetOrderDetailsOutput struct {
	Items []OrderItem
}
