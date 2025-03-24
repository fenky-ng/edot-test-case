package user

import (
	"context"

	"github.com/fenky-ng/edot-test-case/order/internal/model"
)

type RepoHttpUserInterface interface {
	GetUserProfile(ctx context.Context, input model.GetUserProfileInput) (output model.GetUserProfileOutput, err error)
}
