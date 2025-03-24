package usecase

import (
	"github.com/fenky-ng/edot-test-case/order/internal/usecase/order"
	"github.com/fenky-ng/edot-test-case/order/internal/usecase/user"
)

var Auth user.UserUsecaseInterface

type Usecase struct {
	OrderUsecase order.OrderUsecaseInterface
	UserUsecase  user.UserUsecaseInterface
}

type InitUsecaseOptions struct {
	OrderUsecase order.OrderUsecaseInterface
	UserUsecase  user.UserUsecaseInterface
}

func InitUsecase(opts InitUsecaseOptions) *Usecase {
	Auth = opts.UserUsecase

	uc := &Usecase{
		OrderUsecase: opts.OrderUsecase,
		UserUsecase:  opts.UserUsecase,
	}

	return uc
}
