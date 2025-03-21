package shop

import (
	"context"

	"github.com/fenky-ng/edot-test-case/shop/internal/model"
)

type ShopUsecaseInterface interface {
	CreateShop(ctx context.Context, input model.CreateShopInput) (output model.CreateShopOutput, err error)
	GetMyShop(ctx context.Context, input model.GetMyShopInput) (output model.GetMyShopOutput, err error)
	GetShops(ctx context.Context, input model.GetShopsInput) (output model.GetShopsOutput, err error)
	GetShopById(ctx context.Context, input model.GetShopByIdInput) (output model.GetShopByIdOutput, err error)
}
