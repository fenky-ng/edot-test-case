package order

import (
	"context"

	"github.com/fenky-ng/edot-test-case/order/internal/model"
	dbtx "github.com/fenky-ng/edot-test-case/order/internal/utility/database/tx"
)

type RepoDbOrderInterface interface {
	dbtx.DbTxInterface
	InsertOrder(ctx context.Context, input model.InsertOrderInput) (output model.InsertOrderOutput, err error)
	InsertOrderDetails(ctx context.Context, input model.InsertOrderDetailsInput) (output model.InsertOrderDetailsOutput, err error)
	UpdateOrder(ctx context.Context, input model.UpdateOrderInput) (output model.UpdateOrderOutput, err error)
	GetOrders(ctx context.Context, input model.GetOrdersInput) (output model.GetOrdersOutput, err error)
	GetOrderDetails(ctx context.Context, input model.GetOrderDetailsInput) (output model.GetOrderDetailsOutput, err error)
}
