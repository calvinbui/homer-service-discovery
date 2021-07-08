package config

import (
	"fmt"

	"github.com/calvinbui/homer-docker-service-discovery/internal/docker"
	"github.com/calvinbui/homer-docker-service-discovery/internal/homer"
	"github.com/calvinbui/homer-docker-service-discovery/internal/logger"
	"github.com/docker/docker/client"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Docker *client.Client

	LogLevel *string `env:"LOG_LEVEL" envDefault:"Info"`

	HomerConfig     *homer.Config
	HomerConfigPath *string `env:"HOMER_CONFIG" envDefault:"./test/config.yaml"`
}

func New() (*Config, error) {
	var err error
	conf := Config{}

	if err := env.Parse(&conf); err != nil {
		return nil, fmt.Errorf("Error parsing config from env: %+v\n", err)
	}

	conf.Docker, err = docker.CreateClient()
	if err != nil {
		return nil, fmt.Errorf("Error creating Docker client: %w", err)
	}

	err = logger.SetLevel(conf.LogLevel)
	if err != nil {
		return nil, fmt.Errorf("Error setting log level: %w", err)
	}

	conf.HomerConfig, err = homer.GetConfig(*conf.HomerConfigPath)
	if err != nil {
		return nil, fmt.Errorf("Error getting Homer config: %w", err)
	}

	return &conf, nil
}
