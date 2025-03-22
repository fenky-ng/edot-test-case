//go:build wireinject
// +build wireinject

package main

import (
	wire "github.com/google/wire"

	"github.com/fenky-ng/edot-test-case/product/internal/config"
	repo_db_product "github.com/fenky-ng/edot-test-case/product/internal/repository/db/product"
	repo_http_user "github.com/fenky-ng/edot-test-case/product/internal/repository/http/user"
	"github.com/fenky-ng/edot-test-case/product/internal/usecase"
	usecase_product "github.com/fenky-ng/edot-test-case/product/internal/usecase/product"
	usecase_user "github.com/fenky-ng/edot-test-case/product/internal/usecase/user"
	db "github.com/fenky-ng/edot-test-case/product/internal/utility/database"
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

		// product db repository
		wire.Struct(new(repo_db_product.InitRepoDbProductOptions), "*"),
		repo_db_product.InitRepoDbProduct,
		wire.Bind(new(repo_db_product.RepoDbProductInterface), new(*repo_db_product.RepoDbProduct)),

		// product usecase
		wire.Struct(new(usecase_product.InitProductUsecaseOptions), "*"),
		usecase_product.InitProductUsecase,
		wire.Bind(new(usecase_product.ProductUsecaseInterface), new(*usecase_product.ProductUsecase)),

		// usecases
		wire.Struct(new(usecase.InitUsecaseOptions), "*"),
		usecase.InitUsecase,
	)

	return nil, nil
}
