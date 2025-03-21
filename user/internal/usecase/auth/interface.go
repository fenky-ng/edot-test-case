package auth

import (
	"context"

	"github.com/fenky-ng/edot-test-case/user/internal/model"
)

type AuthUsecaseInterface interface {
	Register(ctx context.Context, input model.RegisterInput) (output model.RegisterOutput, err error)
	Login(ctx context.Context, input model.LoginInput) (output model.LoginOutput, err error)
}
