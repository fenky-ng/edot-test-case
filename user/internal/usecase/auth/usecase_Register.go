package auth

import (
	"context"
	"errors"
	"time"

	"github.com/fenky-ng/edot-test-case/user/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	hash_util "github.com/fenky-ng/edot-test-case/user/internal/utility/hash"
	"github.com/google/uuid"
)

func (u *AuthUsecase) Register(ctx context.Context, input model.RegisterInput) (output model.RegisterOutput, err error) {
	userOut, err := u.repoDbUser.GetUser(ctx, model.GetUserInput{
		Phone: input.Phone,
		Email: input.Email,
	})
	if err != nil && !errors.Is(err, in_err.ErrUserNotFound) {
		return output, err
	}

	if userOut.Id != uuid.Nil {
		if input.Phone != "" && userOut.Phone == input.Phone {
			err = in_err.ErrPhoneRegistered
		}
		if input.Email != "" && userOut.Email == input.Email {
			err = in_err.ErrEmailRegistered
		}
		if userOut.DeletedAt != 0 {
			err = in_err.ErrUserDeactivated
		}
		return output, err
	}

	insertOut, err := u.repoDbUser.InsertUser(ctx, model.InsertUserInput{
		Id:             uuid.New(),
		Name:           input.Name,
		Phone:          input.Phone,
		Email:          input.Email,
		HashedPassword: hash_util.HashPassword(input.Password),
		CreatedAt:      time.Now().UnixMilli(),
		CreatedBy:      constant.ServiceName,
	})
	if err != nil {
		return output, err
	}

	output.Id = insertOut.Id

	return output, nil
}
