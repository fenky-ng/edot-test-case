package config

import (
	"github.com/caarlos0/env/v10"
)

func Init() (cfg *Configuration, err error) {
	cfg = new(Configuration)

	err = env.Parse(cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
