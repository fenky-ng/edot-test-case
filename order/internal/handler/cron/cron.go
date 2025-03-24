package cron

import (
	"github.com/fenky-ng/edot-test-case/order/internal/config"
	"github.com/fenky-ng/edot-test-case/order/internal/usecase"
)

type CronHandler struct {
	config  *config.Configuration
	usecase *usecase.Usecase
}

type InitCronHandlerCOptions struct {
	Config  *config.Configuration
	Usecase *usecase.Usecase
}

func InitCronHandler(opts InitCronHandlerCOptions) *CronHandler {
	return &CronHandler{
		config:  opts.Config,
		usecase: opts.Usecase,
	}
}
