package config

import (
	"fmt"
	"github.com/calvinbui/homer-docker-service-discovery/internal/consul"
	"github.com/calvinbui/homer-docker-service-discovery/internal/docker"
	"github.com/calvinbui/homer-docker-service-discovery/internal/logger"
	"github.com/docker/docker/client"
	"github.com/hashicorp/consul/api"

	"github.com/caarlos0/env/v6"
)

const (
	Docker = "Docker"
	Consul = "Consul"
)

type Config struct {
	Docker *client.Client

	Consul *api.Client

	ServiceDiscovery string `env:"SERVICE_DISCOVERY" envDefault:"Docker"`

	ConsulHost string `env:"CONSUL_HOST" envDefault:"127.0.0.1:8500"`

	LogLevel string `env:"LOG_LEVEL" envDefault:"Info"`

	HomerBaseConfigPath string `env:"HOMER_BASE_CONFIG" envDefault:"/base.yml"`

	HomerConfigPath string `env:"HOMER_CONFIG" envDefault:"/config.yml"`
}

func New() (Config, error) {
	var err error
	conf := Config{}

	if err := env.Parse(&conf); err != nil {
		return Config{}, fmt.Errorf("Error parsing config from env: %+v\n", err)
	}

	if conf.ServiceDiscovery == Docker {
		conf.Docker, err = docker.CreateClient()
		if err != nil {
			return Config{}, fmt.Errorf("Error creating Docker client: %w", err)
		}
	} else if conf.ServiceDiscovery == Consul {
		conf.Consul, err = consul.CreateClient(conf.ConsulHost)
		if err != nil {
			return Config{}, fmt.Errorf("Error creating Consul client: %w", err)
		}
	} else {
		return Config{}, fmt.Errorf("Unknow Service Discovery in configuration")
	}
	err = logger.SetLevel(conf.LogLevel)
	if err != nil {
		return Config{}, fmt.Errorf("Error setting log level: %w", err)
	}

	logger.Debug(fmt.Sprintf("%+v", conf))

	return conf, nil
}
