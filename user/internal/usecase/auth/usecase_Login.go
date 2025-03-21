package auth

import (
	"context"
	"errors"

	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	hash_util "github.com/fenky-ng/edot-test-case/user/internal/utility/hash"
	"github.com/fenky-ng/edot-test-case/user/internal/utility/jwt"
)

func (u *AuthUsecase) Login(ctx context.Context, input model.LoginInput) (output model.LoginOutput, err error) {
	retErr := in_err.ErrInvalidPhoneLogin
	if input.Phone == "" {
		retErr = in_err.ErrInvalidEmailLogin
	}

	userOut, err := u.repoDbUser.GetUser(ctx, model.GetUserInput{
		Phone: input.Phone,
		Email: input.Email,
	})
	if err != nil {
		if errors.Is(err, in_err.ErrUserNotFound) {
			err = errors.Join(retErr, err)
		}
		return output, err
	}

	if userOut.HashedPassword != hash_util.HashPassword(input.Password) {
		err = errors.Join(retErr, err)
		return output, err
	}

	if userOut.DeletedAt != 0 {
		err = in_err.ErrUserDeactivated
		return output, err
	}

	jwt, err := jwt.GenerateJWT(userOut.Id)
	if err != nil {
		err = errors.Join(in_err.ErrJWT, err)
		return output, err
	}

	output.JWT = jwt

	return output, nil
}
