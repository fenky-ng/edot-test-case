package shop

import (
	"context"

	in_err "github.com/fenky-ng/edot-test-case/shop/internal/error"
	"github.com/fenky-ng/edot-test-case/shop/internal/model"
)

func (u *ShopUsecase) GetMyShop(ctx context.Context, input model.GetMyShopInput) (output model.GetMyShopOutput, err error) {
	shopOut, err := u.repoDbShop.GetShop(ctx, model.GetShopInput{
		OwnerId: input.UserId,
	})
	if err != nil {
		return output, err
	}

	if shopOut.DeletedAt != 0 {
		err = in_err.ErrShopDeactivated
		return output, err
	}

	output.Id = shopOut.Id
	output.OwnerId = shopOut.OwnerId
	output.Name = shopOut.Name
	output.Status = shopOut.Status

	return output, nil
}
