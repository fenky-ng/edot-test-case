package shop

import "github.com/fenky-ng/edot-test-case/order/internal/config"

type RepoHttpShop struct {
	config *config.Configuration
}

type InitRepoHttpShopOptions struct {
	Config *config.Configuration
}

func InitRepoHttpShop(opts InitRepoHttpShopOptions) *RepoHttpShop {
	return &RepoHttpShop{
		config: opts.Config,
	}
}
