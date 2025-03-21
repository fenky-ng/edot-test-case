package auth

import (
	db_user "github.com/fenky-ng/edot-test-case/user/internal/repository/db/user"
)

type AuthUsecase struct {
	repoDbUser db_user.RepoDbUserInterface
}

type InitAuthUsecaseOptions struct {
	RepoDbUser db_user.RepoDbUserInterface
}

func InitAuthUsecase(opts InitAuthUsecaseOptions) *AuthUsecase {
	return &AuthUsecase{
		repoDbUser: opts.RepoDbUser,
	}
}
