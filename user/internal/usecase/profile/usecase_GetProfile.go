package profile

import (
	"context"

	"github.com/fenky-ng/edot-test-case/user/internal/model"
)

func (u *ProfileUsecase) GetProfile(ctx context.Context, input model.GetProfileInput) (output model.GetProfileOutput, err error) {
	userOut, err := u.repoDbUser.GetUser(ctx, model.GetUserInput{
		Id: input.Id,
	})
	if err != nil {
		return output, err
	}

	output.Id = userOut.Id
	output.Name = userOut.Name
	output.Phone = userOut.Phone
	output.Email = userOut.Email

	return output, nil
}
