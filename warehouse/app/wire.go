//go:build wireinject
// +build wireinject

package main

import (
	wire "github.com/google/wire"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/config"
	repo_db_warehouse "github.com/fenky-ng/edot-test-case/warehouse/internal/repository/db/warehouse"
	repo_http_product "github.com/fenky-ng/edot-test-case/warehouse/internal/repository/http/product"
	repo_http_shop "github.com/fenky-ng/edot-test-case/warehouse/internal/repository/http/shop"
	repo_http_user "github.com/fenky-ng/edot-test-case/warehouse/internal/repository/http/user"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/usecase"
	usecase_user "github.com/fenky-ng/edot-test-case/warehouse/internal/usecase/user"
	usecase_warehouse "github.com/fenky-ng/edot-test-case/warehouse/internal/usecase/warehouse"
	db "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/database"
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

		// warehouse db repository
		wire.Struct(new(repo_db_warehouse.InitRepoDbWarehouseOptions), "*"),
		repo_db_warehouse.InitRepoDbWarehouse,
		wire.Bind(new(repo_db_warehouse.RepoDbWarehouseInterface), new(*repo_db_warehouse.RepoDbWarehouse)),

		// shop http repository
		wire.Struct(new(repo_http_shop.InitRepoHttpShopOptions), "*"),
		repo_http_shop.InitRepoHttpShop,
		wire.Bind(new(repo_http_shop.RepoHttpShopInterface), new(*repo_http_shop.RepoHttpShop)),

		// product http repository
		wire.Struct(new(repo_http_product.InitRepoHttpProductOptions), "*"),
		repo_http_product.InitRepoHttpProduct,
		wire.Bind(new(repo_http_product.RepoHttpProductInterface), new(*repo_http_product.RepoHttpProduct)),

		// warehouse usecase
		wire.Struct(new(usecase_warehouse.InitWarehouseUsecaseOptions), "*"),
		usecase_warehouse.InitWarehouseUsecase,
		wire.Bind(new(usecase_warehouse.WarehouseUsecaseInterface), new(*usecase_warehouse.WarehouseUsecase)),

		// usecases
		wire.Struct(new(usecase.InitUsecaseOptions), "*"),
		usecase.InitUsecase,
	)

	return nil, nil
}
