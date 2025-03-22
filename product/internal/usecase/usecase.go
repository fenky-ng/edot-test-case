package usecase

import (
	"github.com/fenky-ng/edot-test-case/product/internal/usecase/product"
	"github.com/fenky-ng/edot-test-case/product/internal/usecase/user"
)

var Auth user.UserUsecaseInterface

type Usecase struct {
	ProductUsecase product.ProductUsecaseInterface
	UserUsecase    user.UserUsecaseInterface
}

type InitUsecaseOptions struct {
	ProductUsecase product.ProductUsecaseInterface
	UserUsecase    user.UserUsecaseInterface
}

func InitUsecase(opts InitUsecaseOptions) *Usecase {
	Auth = opts.UserUsecase

	uc := &Usecase{
		ProductUsecase: opts.ProductUsecase,
		UserUsecase:    opts.UserUsecase,
	}

	return uc
}
