package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	Env string `env:"example" envDefault:"default"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return &Config{}, nil
}
