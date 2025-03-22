package product

import (
	db_product "github.com/fenky-ng/edot-test-case/product/internal/repository/db/product"
	http_shop "github.com/fenky-ng/edot-test-case/product/internal/repository/http/shop"
	http_warehouse "github.com/fenky-ng/edot-test-case/product/internal/repository/http/warehouse"
)

type ProductUsecase struct {
	repoDbProduct     db_product.RepoDbProductInterface
	repoHttpShop      http_shop.RepoHttpShopInterface
	repoHttpWarehouse http_warehouse.RepoHttpWarehouseInterface
}

type InitProductUsecaseOptions struct {
	RepoDbProduct     db_product.RepoDbProductInterface
	RepoHttpShop      http_shop.RepoHttpShopInterface
	RepoHttpWarehouse http_warehouse.RepoHttpWarehouseInterface
}

func InitProductUsecase(opts InitProductUsecaseOptions) *ProductUsecase {
	return &ProductUsecase{
		repoDbProduct:     opts.RepoDbProduct,
		repoHttpShop:      opts.RepoHttpShop,
		repoHttpWarehouse: opts.RepoHttpWarehouse,
	}
}
