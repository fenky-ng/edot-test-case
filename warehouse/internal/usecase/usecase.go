package usecase

import (
	"github.com/fenky-ng/edot-test-case/warehouse/internal/usecase/user"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/usecase/warehouse"
)

var Auth user.UserUsecaseInterface

type Usecase struct {
	WarehouseUsecase warehouse.WarehouseUsecaseInterface
	UserUsecase      user.UserUsecaseInterface
}

type InitUsecaseOptions struct {
	WarehouseUsecase warehouse.WarehouseUsecaseInterface
	UserUsecase      user.UserUsecaseInterface
}

func InitUsecase(opts InitUsecaseOptions) *Usecase {
	Auth = opts.UserUsecase

	uc := &Usecase{
		WarehouseUsecase: opts.WarehouseUsecase,
		UserUsecase:      opts.UserUsecase,
	}

	return uc
}
