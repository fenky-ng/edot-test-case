package warehouse

import (
	db_warehouse "github.com/fenky-ng/edot-test-case/warehouse/internal/repository/db/warehouse"
)

type WarehouseUsecase struct {
	repoDbWarehouse db_warehouse.RepoDbWarehouseInterface
}

type InitWarehouseUsecaseOptions struct {
	RepoDbWarehouse db_warehouse.RepoDbWarehouseInterface
}

func InitWarehouseUsecase(opts InitWarehouseUsecaseOptions) *WarehouseUsecase {
	return &WarehouseUsecase{
		repoDbWarehouse: opts.RepoDbWarehouse,
	}
}
