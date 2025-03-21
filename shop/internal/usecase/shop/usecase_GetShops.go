package shop

import (
	"context"

	"github.com/fenky-ng/edot-test-case/shop/internal/model"
)

func (u *ShopUsecase) GetShops(ctx context.Context, input model.GetShopsInput) (output model.GetShopsOutput, err error) {
	return u.repoDbShop.GetShops(ctx, input)
}
