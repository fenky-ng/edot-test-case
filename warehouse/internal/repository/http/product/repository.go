package product

import "github.com/fenky-ng/edot-test-case/warehouse/internal/config"

type RepoHttpProduct struct {
	config *config.Configuration
}

type InitRepoHttpProductOptions struct {
	Config *config.Configuration
}

func InitRepoHttpProduct(opts InitRepoHttpProductOptions) *RepoHttpProduct {
	return &RepoHttpProduct{
		config: opts.Config,
	}
}
