package rest

import (
	"github.com/fenky-ng/edot-test-case/warehouse/internal/config"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/usecase"
)

type RestAPI struct {
	config  *config.Configuration
	usecase *usecase.Usecase
}

type InitRestAPIOptions struct {
	Config  *config.Configuration
	Usecase *usecase.Usecase
}

func InitRestAPI(opts InitRestAPIOptions) *RestAPI {
	return &RestAPI{
		config:  opts.Config,
		usecase: opts.Usecase,
	}
}
