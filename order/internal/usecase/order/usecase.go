package order

import (
	db_order "github.com/fenky-ng/edot-test-case/order/internal/repository/db/order"
	http_product "github.com/fenky-ng/edot-test-case/order/internal/repository/http/product"
	http_shop "github.com/fenky-ng/edot-test-case/order/internal/repository/http/shop"
	http_warehouse "github.com/fenky-ng/edot-test-case/order/internal/repository/http/warehouse"
)

type OrderUsecase struct {
	repoDbOrder       db_order.RepoDbOrderInterface
	repoHttpShop      http_shop.RepoHttpShopInterface
	repoHttpProduct   http_product.RepoHttpProductInterface
	repoHttpWarehouse http_warehouse.RepoHttpWarehouseInterface

	// self interface
	OrderUsecase OrderUsecaseInterface
}

type InitOrderUsecaseOptions struct {
	RepoDbOrder       db_order.RepoDbOrderInterface
	RepoHttpShop      http_shop.RepoHttpShopInterface
	RepoHttpProduct   http_product.RepoHttpProductInterface
	RepoHttpWarehouse http_warehouse.RepoHttpWarehouseInterface
}

func InitOrderUsecase(opts InitOrderUsecaseOptions) *OrderUsecase {
	u := &OrderUsecase{
		repoDbOrder:       opts.RepoDbOrder,
		repoHttpShop:      opts.RepoHttpShop,
		repoHttpProduct:   opts.RepoHttpProduct,
		repoHttpWarehouse: opts.RepoHttpWarehouse,
	}
	u.OrderUsecase = u
	return u
}
