package profile

import (
	"context"

	"github.com/fenky-ng/edot-test-case/user/internal/model"
)

type ProfileUsecaseInterface interface {
	GetProfile(ctx context.Context, input model.GetProfileInput) (output model.GetProfileOutput, err error)
}
