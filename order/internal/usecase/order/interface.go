package order

import (
	"context"

	"github.com/fenky-ng/edot-test-case/order/internal/model"
)

type OrderUsecaseInterface interface {
	CreateOrder(ctx context.Context, input model.CreateOrderInput) (output model.CreateOrderOutput, err error)
	UpdateOrder(ctx context.Context, input model.UpdateOrderInput) (output model.UpdateOrderOutput, err error)
	GetOrders(ctx context.Context, input model.GetOrdersInput) (output model.GetOrdersOutput, err error)
	ReleaseStocks(ctx context.Context, input model.ReleaseStocksInput) (output model.ReleaseStocksOutput, err error)

	// INTERNAL - order usecase related functions
	ValidateCreateOrder(ctx context.Context, input model.CreateOrderInput) (err error)
}
