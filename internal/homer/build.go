package homer

import "github.com/calvinbui/homer-docker-service-discovery/internal/docker"

func BuildConfig(config Config, containers []docker.Container) (Config, error) {
	return config, nil
}
