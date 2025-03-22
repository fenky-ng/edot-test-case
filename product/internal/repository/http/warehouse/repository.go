package warehouse

import "github.com/fenky-ng/edot-test-case/product/internal/config"

type RepoHttpWarehouse struct {
	config *config.Configuration
}

type InitRepoHttpWarehouseOptions struct {
	Config *config.Configuration
}

func InitRepoHttpWarehouse(opts InitRepoHttpWarehouseOptions) *RepoHttpWarehouse {
	return &RepoHttpWarehouse{
		config: opts.Config,
	}
}
