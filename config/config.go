package config

import (
	"github.com/caarlos0/env/v10"
)

type POW struct {
	Complexity uint8 `env:"POW_COMPLEXITY"`
}

type Common struct {
	POW
	LogLevel string `env:"LOG_LEVEL"`
}

type ServerConfig struct {
	Common
}

type ClientConfig struct {
	ServerAddr string `env:"SERVER_ADDR"`
	Common
}

func NewClientConfig() (*ClientConfig, error) {
	cfg := &ClientConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewServerConfig() (*ServerConfig, error) {
	cfg := &ServerConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
