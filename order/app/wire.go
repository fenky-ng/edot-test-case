//go:build wireinject
// +build wireinject

package main

import (
	wire "github.com/google/wire"

	"github.com/fenky-ng/edot-test-case/order/internal/config"
	repo_db_order "github.com/fenky-ng/edot-test-case/order/internal/repository/db/order"
	repo_http_product "github.com/fenky-ng/edot-test-case/order/internal/repository/http/product"
	repo_http_shop "github.com/fenky-ng/edot-test-case/order/internal/repository/http/shop"
	repo_http_user "github.com/fenky-ng/edot-test-case/order/internal/repository/http/user"
	repo_http_warehouse "github.com/fenky-ng/edot-test-case/order/internal/repository/http/warehouse"
	"github.com/fenky-ng/edot-test-case/order/internal/usecase"
	usecase_order "github.com/fenky-ng/edot-test-case/order/internal/usecase/order"
	usecase_user "github.com/fenky-ng/edot-test-case/order/internal/usecase/user"
	db "github.com/fenky-ng/edot-test-case/order/internal/utility/database"
	dbtx "github.com/fenky-ng/edot-test-case/order/internal/utility/database/tx"
)

func InitializeUsecase(
	cfg *config.Configuration,
) (*usecase.Usecase, error) {
	wire.Build(
		// db
		db.InitDatabase,

		// dbtx
		wire.Struct(new(dbtx.InitDbTxOptions), "*"),
		dbtx.InitDbTx,
		wire.Bind(new(dbtx.DbTxInterface), new(*dbtx.DbTx)),

		// user http repository
		wire.Struct(new(repo_http_user.InitRepoHttpUserOptions), "*"),
		repo_http_user.InitRepoHttpUser,
		wire.Bind(new(repo_http_user.RepoHttpUserInterface), new(*repo_http_user.RepoHttpUser)),

		// user usecase
		wire.Struct(new(usecase_user.InitUserUsecaseOptions), "*"),
		usecase_user.InitUserUsecase,
		wire.Bind(new(usecase_user.UserUsecaseInterface), new(*usecase_user.UserUsecase)),

		// order db repository
		wire.Struct(new(repo_db_order.InitRepoDbOrderOptions), "*"),
		repo_db_order.InitRepoDbOrder,
		wire.Bind(new(repo_db_order.RepoDbOrderInterface), new(*repo_db_order.RepoDbOrder)),

		// shop http repository
		wire.Struct(new(repo_http_shop.InitRepoHttpShopOptions), "*"),
		repo_http_shop.InitRepoHttpShop,
		wire.Bind(new(repo_http_shop.RepoHttpShopInterface), new(*repo_http_shop.RepoHttpShop)),

		// product http repository
		wire.Struct(new(repo_http_product.InitRepoHttpProductOptions), "*"),
		repo_http_product.InitRepoHttpProduct,
		wire.Bind(new(repo_http_product.RepoHttpProductInterface), new(*repo_http_product.RepoHttpProduct)),

		// warehouse http repository
		wire.Struct(new(repo_http_warehouse.InitRepoHttpWarehouseOptions), "*"),
		repo_http_warehouse.InitRepoHttpWarehouse,
		wire.Bind(new(repo_http_warehouse.RepoHttpWarehouseInterface), new(*repo_http_warehouse.RepoHttpWarehouse)),

		// order usecase
		wire.Struct(new(usecase_order.InitOrderUsecaseOptions), "*"),
		usecase_order.InitOrderUsecase,
		wire.Bind(new(usecase_order.OrderUsecaseInterface), new(*usecase_order.OrderUsecase)),

		// usecases
		wire.Struct(new(usecase.InitUsecaseOptions), "*"),
		usecase.InitUsecase,
	)

	return nil, nil
}
