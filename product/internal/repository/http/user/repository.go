package user

import "github.com/fenky-ng/edot-test-case/product/internal/config"

type RepoHttpUser struct {
	config *config.Configuration
}

type InitRepoHttpUserOptions struct {
	Config *config.Configuration
}

func InitRepoHttpUser(opts InitRepoHttpUserOptions) *RepoHttpUser {
	return &RepoHttpUser{
		config: opts.Config,
	}
}
