//go:build wireinject
// +build wireinject

package main

import (
	wire "github.com/google/wire"

	"github.com/fenky-ng/edot-test-case/shop/internal/config"
	repo_db_shop "github.com/fenky-ng/edot-test-case/shop/internal/repository/db/shop"
	repo_http_user "github.com/fenky-ng/edot-test-case/shop/internal/repository/http/user"
	"github.com/fenky-ng/edot-test-case/shop/internal/usecase"
	usecase_shop "github.com/fenky-ng/edot-test-case/shop/internal/usecase/shop"
	usecase_user "github.com/fenky-ng/edot-test-case/shop/internal/usecase/user"
	db "github.com/fenky-ng/edot-test-case/shop/internal/utility/database"
)

func InitializeUsecase(
	cfg *config.Configuration,
) (*usecase.Usecase, error) {
	wire.Build(
		// db
		db.InitDatabase,

		// user http repository
		wire.Struct(new(repo_http_user.InitRepoHttpUserOptions), "*"),
		repo_http_user.InitRepoHttpUser,
		wire.Bind(new(repo_http_user.RepoHttpUserInterface), new(*repo_http_user.RepoHttpUser)),

		// user usecase
		wire.Struct(new(usecase_user.InitUserUsecaseOptions), "*"),
		usecase_user.InitUserUsecase,
		wire.Bind(new(usecase_user.UserUsecaseInterface), new(*usecase_user.UserUsecase)),

		// shop db repository
		wire.Struct(new(repo_db_shop.InitRepoDbShopOptions), "*"),
		repo_db_shop.InitRepoDbShop,
		wire.Bind(new(repo_db_shop.RepoDbShopInterface), new(*repo_db_shop.RepoDbShop)),

		// shop usecase
		wire.Struct(new(usecase_shop.InitShopUsecaseOptions), "*"),
		usecase_shop.InitShopUsecase,
		wire.Bind(new(usecase_shop.ShopUsecaseInterface), new(*usecase_shop.ShopUsecase)),

		// usecases
		wire.Struct(new(usecase.InitUsecaseOptions), "*"),
		usecase.InitUsecase,
	)

	return nil, nil
}
