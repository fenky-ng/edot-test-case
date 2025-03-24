package shop

import (
	"context"

	"github.com/fenky-ng/edot-test-case/order/internal/model"
)

type RepoHttpShopInterface interface {
	GetMyShop(ctx context.Context, input model.GetMyShopInput) (output model.GetMyShopOutput, err error)
}
