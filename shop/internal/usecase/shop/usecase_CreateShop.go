package shop

import (
	"context"
	"errors"
	"time"

	"github.com/fenky-ng/edot-test-case/shop/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/shop/internal/error"
	"github.com/fenky-ng/edot-test-case/shop/internal/model"
	"github.com/google/uuid"
)

func (u *ShopUsecase) CreateShop(ctx context.Context, input model.CreateShopInput) (output model.CreateShopOutput, err error) {
	shopOut, err := u.repoDbShop.GetShop(ctx, model.GetShopInput{
		OwnerId: input.UserId,
	})
	if err != nil && !errors.Is(err, in_err.ErrShopNotFound) {
		return output, err
	}

	if shopOut.DeletedAt != 0 {
		err = in_err.ErrShopDeactivated
		return output, err
	} else if shopOut.Id != uuid.Nil {
		err = in_err.ErrAlreadyOwnAShop
		return output, err
	}

	insertOut, err := u.repoDbShop.InsertShop(ctx, model.InsertShopInput{
		Id:        uuid.New(),
		OwnerId:   input.UserId,
		Name:      input.Name,
		Status:    constant.ShopStatus_Active,
		CreatedAt: time.Now().UnixMilli(),
		CreatedBy: input.UserId,
	})
	if err != nil {
		return output, err
	}

	output.Id = insertOut.Id

	return output, nil
}
