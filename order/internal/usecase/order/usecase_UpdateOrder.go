package order

import (
	"context"
	"time"

	in_err "github.com/fenky-ng/edot-test-case/order/internal/error"
	"github.com/fenky-ng/edot-test-case/order/internal/model"
)

func (u *OrderUsecase) UpdateOrder(ctx context.Context, input model.UpdateOrderInput) (output model.UpdateOrderOutput, err error) {
	ordersOut, err := u.repoDbOrder.GetOrders(ctx, model.GetOrdersInput{
		OrderNo: input.OrderNo,
	})
	if err != nil {
		return output, err
	}
	if len(ordersOut.Orders) == 0 {
		err = in_err.ErrOrderNotFound
		return output, err
	}

	input.UpdatedAt = time.Now().UnixMilli()
	return u.repoDbOrder.UpdateOrder(ctx, input)
}
