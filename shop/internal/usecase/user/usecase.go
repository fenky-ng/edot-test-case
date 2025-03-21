package user

import (
	http_user "github.com/fenky-ng/edot-test-case/shop/internal/repository/http/user"
)

type UserUsecase struct {
	repoHttpUser http_user.RepoHttpUserInterface
}

type InitUserUsecaseOptions struct {
	RepoHttpUser http_user.RepoHttpUserInterface
}

func InitUserUsecase(opts InitUserUsecaseOptions) *UserUsecase {
	return &UserUsecase{
		repoHttpUser: opts.RepoHttpUser,
	}
}
