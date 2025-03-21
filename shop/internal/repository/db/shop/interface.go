package shop

import (
	"context"

	"github.com/fenky-ng/edot-test-case/shop/internal/model"
)

type RepoDbShopInterface interface {
	GetShops(ctx context.Context, input model.GetShopsInput) (output model.GetShopsOutput, err error)
	GetShop(ctx context.Context, input model.GetShopInput) (output model.GetShopOutput, err error)
	InsertShop(ctx context.Context, input model.InsertShopInput) (output model.InsertShopOutput, err error)
}
