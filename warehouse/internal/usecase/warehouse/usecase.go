package warehouse

import (
	db_warehouse "github.com/fenky-ng/edot-test-case/warehouse/internal/repository/db/warehouse"
	http_product "github.com/fenky-ng/edot-test-case/warehouse/internal/repository/http/product"
	http_shop "github.com/fenky-ng/edot-test-case/warehouse/internal/repository/http/shop"
)

type WarehouseUsecase struct {
	repoDbWarehouse db_warehouse.RepoDbWarehouseInterface
	repoHttpShop    http_shop.RepoHttpShopInterface
	repoHttpProduct http_product.RepoHttpProductInterface

	// self
	warehouseUsecase WarehouseUsecaseInterface
}

type InitWarehouseUsecaseOptions struct {
	RepoDbWarehouse db_warehouse.RepoDbWarehouseInterface
	RepoHttpShop    http_shop.RepoHttpShopInterface
	RepoHttpProduct http_product.RepoHttpProductInterface
}

func InitWarehouseUsecase(opts InitWarehouseUsecaseOptions) *WarehouseUsecase {
	u := &WarehouseUsecase{
		repoDbWarehouse: opts.RepoDbWarehouse,
		repoHttpShop:    opts.RepoHttpShop,
		repoHttpProduct: opts.RepoHttpProduct,
	}
	u.warehouseUsecase = u
	return u
}
