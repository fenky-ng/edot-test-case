package usecase

import (
	"github.com/fenky-ng/edot-test-case/shop/internal/usecase/shop"
	"github.com/fenky-ng/edot-test-case/shop/internal/usecase/user"
)

var Auth user.UserUsecaseInterface

type Usecase struct {
	ShopUsecase shop.ShopUsecaseInterface
	UserUsecase user.UserUsecaseInterface
}

type InitUsecaseOptions struct {
	ShopUsecase shop.ShopUsecaseInterface
	UserUsecase user.UserUsecaseInterface
}

func InitUsecase(opts InitUsecaseOptions) *Usecase {
	Auth = opts.UserUsecase

	uc := &Usecase{
		ShopUsecase: opts.ShopUsecase,
		UserUsecase: opts.UserUsecase,
	}

	return uc
}
