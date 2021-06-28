package config

import (
	"fmt"

	"github.com/calvinbui/homer-docker-service-discovery/internal/docker"
	"github.com/calvinbui/homer-docker-service-discovery/internal/logger"
	"github.com/docker/docker/client"

	"github.com/caarlos0/env/v6"
)

type config struct {
	Docker *client.Client

	LogLevel *string `env:"LOG_LEVEL" envDefault:"Info"`
}

func New() (*config, error) {
	var err error
	conf := config{}

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

	return &conf, nil
}
