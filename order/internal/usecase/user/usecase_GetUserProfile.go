package user

import (
	"context"

	"github.com/fenky-ng/edot-test-case/order/internal/model"
)

func (s *UserUsecase) GetUserProfile(ctx context.Context, input model.GetUserProfileInput) (output model.GetUserProfileOutput, err error) {
	return s.repoHttpUser.GetUserProfile(ctx, input)
}
