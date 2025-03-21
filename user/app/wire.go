//go:build wireinject
// +build wireinject

package main

import (
	wire "github.com/google/wire"

	"github.com/fenky-ng/edot-test-case/user/internal/config"
	repo_db_user "github.com/fenky-ng/edot-test-case/user/internal/repository/db/user"
	"github.com/fenky-ng/edot-test-case/user/internal/usecase"
	usecase_auth "github.com/fenky-ng/edot-test-case/user/internal/usecase/auth"
	usecase_profile "github.com/fenky-ng/edot-test-case/user/internal/usecase/profile"
	db "github.com/fenky-ng/edot-test-case/user/internal/utility/database"
)

func InitializeUsecase(
	cfg *config.Configuration,
) (*usecase.Usecase, error) {
	wire.Build(
		// db
		db.InitDatabase,

		// user db repository
		wire.Struct(new(repo_db_user.InitRepoDbUserOptions), "*"),
		repo_db_user.InitRepoDbUser,
		wire.Bind(new(repo_db_user.RepoDbUserInterface), new(*repo_db_user.RepoDbUser)),

		// auth usecase
		wire.Struct(new(usecase_auth.InitAuthUsecaseOptions), "*"),
		usecase_auth.InitAuthUsecase,
		wire.Bind(new(usecase_auth.AuthUsecaseInterface), new(*usecase_auth.AuthUsecase)),

		// profile usecase
		wire.Struct(new(usecase_profile.InitProfileUsecaseOptions), "*"),
		usecase_profile.InitProfileUsecase,
		wire.Bind(new(usecase_profile.ProfileUsecaseInterface), new(*usecase_profile.ProfileUsecase)),

		// usecases
		wire.Struct(new(usecase.InitUsecaseOptions), "*"),
		usecase.InitUsecase,
	)

	return nil, nil
}
