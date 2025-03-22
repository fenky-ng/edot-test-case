package main

import (
	"github.com/fenky-ng/edot-test-case/product/internal/config"
	"github.com/fenky-ng/edot-test-case/product/internal/usecase"
)

func Init() (
	cfg *config.Configuration,
	usecase *usecase.Usecase,
	err error,
) {
	cfg, err = config.Init()
	if err != nil {
		return
	}

	usecase, err = InitializeUsecase(cfg)
	if err != nil {
		return
	}

	return
}
