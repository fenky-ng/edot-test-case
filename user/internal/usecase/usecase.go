package usecase

import (
	"github.com/fenky-ng/edot-test-case/user/internal/usecase/auth"
	"github.com/fenky-ng/edot-test-case/user/internal/usecase/profile"
)

type Usecase struct {
	AuthUsecase    auth.AuthUsecaseInterface
	ProfileUsecase profile.ProfileUsecaseInterface
}

type InitUsecaseOptions struct {
	AuthUsecase    auth.AuthUsecaseInterface
	ProfileUsecase profile.ProfileUsecaseInterface
}

func InitUsecase(opts InitUsecaseOptions) *Usecase {
	return &Usecase{
		AuthUsecase:    opts.AuthUsecase,
		ProfileUsecase: opts.ProfileUsecase,
	}
}
