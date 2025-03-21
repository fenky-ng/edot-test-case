package shop

import (
	"context"

	"github.com/fenky-ng/edot-test-case/shop/internal/model"
)

func (u *ShopUsecase) GetShopById(ctx context.Context, input model.GetShopByIdInput) (output model.GetShopByIdOutput, err error) {
	shopOut, err := u.repoDbShop.GetShop(ctx, model.GetShopInput{
		ShopId:         input.Id,
		ExcludeDeleted: true,
	})
	if err != nil {
		return output, err
	}

	output.Id = shopOut.Id
	output.OwnerId = shopOut.OwnerId
	output.Name = shopOut.Name
	output.Status = shopOut.Status

	return output, nil
}
