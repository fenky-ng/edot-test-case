package shop

import (
	"context"

	"github.com/fenky-ng/edot-test-case/product/internal/model"
)

type RepoHttpShopInterface interface {
	GetMyShop(ctx context.Context, input model.GetMyShopInput) (output model.GetMyShopOutput, err error)
	GetShops(ctx context.Context, input model.GetShopsInput) (output model.GetShopsOutput, err error)
}
