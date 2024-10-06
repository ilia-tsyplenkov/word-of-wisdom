package config

import (
	"github.com/caarlos0/env/v10"
)

type POW struct {
	Complexity uint8 `env:"POW_COMPLEXITY" envDefault:"10"`
}

type Common struct {
	POW
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
}

type ServerConfig struct {
	Addr           string `env:"SERVER_ADDR" envDefault:"127.0.0.1:8080"`
	RequestLimiter int    `env:"REQUEST_LIMITER" envDefault:"100"`
	Common
}

type ClientConfig struct {
	ServerAddr string `env:"SERVER_ADDR" envDefault:"127.0.0.1:8080"`
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
