package product

import (
	db_product "github.com/fenky-ng/edot-test-case/product/internal/repository/db/product"
)

type ProductUsecase struct {
	repoDbProduct db_product.RepoDbProductInterface
}

type InitProductUsecaseOptions struct {
	RepoDbProduct db_product.RepoDbProductInterface
}

func InitProductUsecase(opts InitProductUsecaseOptions) *ProductUsecase {
	return &ProductUsecase{
		repoDbProduct: opts.RepoDbProduct,
	}
}
