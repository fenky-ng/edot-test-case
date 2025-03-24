package order

import (
	"context"

	"github.com/fenky-ng/edot-test-case/order/internal/model"
)

func (u *OrderUsecase) GetOrders(ctx context.Context, input model.GetOrdersInput) (output model.GetOrdersOutput, err error) {
	// room for improvement: add logic to get order details
	return u.repoDbOrder.GetOrders(ctx, input)
}
