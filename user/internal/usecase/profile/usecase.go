package profile

import (
	db_user "github.com/fenky-ng/edot-test-case/user/internal/repository/db/user"
)

type ProfileUsecase struct {
	repoDbUser db_user.RepoDbUserInterface
}

type InitProfileUsecaseOptions struct {
	RepoDbUser db_user.RepoDbUserInterface
}

func InitProfileUsecase(opts InitProfileUsecaseOptions) *ProfileUsecase {
	return &ProfileUsecase{
		repoDbUser: opts.RepoDbUser,
	}
}
