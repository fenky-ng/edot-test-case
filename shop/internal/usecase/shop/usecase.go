package shop

import (
	db_shop "github.com/fenky-ng/edot-test-case/shop/internal/repository/db/shop"
)

type ShopUsecase struct {
	repoDbShop db_shop.RepoDbShopInterface
}

type InitShopUsecaseOptions struct {
	RepoDbShop db_shop.RepoDbShopInterface
}

func InitShopUsecase(opts InitShopUsecaseOptions) *ShopUsecase {
	return &ShopUsecase{
		repoDbShop: opts.RepoDbShop,
	}
}
