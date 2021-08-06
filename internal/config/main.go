package config

import (
	"fmt"

	"github.com/calvinbui/homer-docker-service-discovery/internal/docker"
	"github.com/calvinbui/homer-docker-service-discovery/internal/logger"
	"github.com/calvinbui/homer-docker-service-discovery/pkg/homer"
	"github.com/docker/docker/client"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Docker *client.Client

	LogLevel string `env:"LOG_LEVEL" envDefault:"Info"`

	HomerBaseConfig     homer.Config
	HomerBaseConfigPath string `env:"HOMER_BASE_CONFIG" envDefault:"../test/base.yml"`

	HomerConfigPath string `env:"HOMER_CONFIG" envDefault:"../test/homer.yml"`
}

func New() (Config, error) {
	var err error
	conf := Config{}

	if err := env.Parse(&conf); err != nil {
		return Config{}, fmt.Errorf("Error parsing config from env: %+v\n", err)
	}

	conf.Docker, err = docker.CreateClient()
	if err != nil {
		return Config{}, fmt.Errorf("Error creating Docker client: %w", err)
	}

	err = logger.SetLevel(conf.LogLevel)
	if err != nil {
		return Config{}, fmt.Errorf("Error setting log level: %w", err)
	}

	conf.HomerBaseConfig, err = homer.GetConfig(conf.HomerBaseConfigPath)
	if err != nil {
		return Config{}, fmt.Errorf("Error getting Homer config: %w", err)
	}

	return conf, nil
}
