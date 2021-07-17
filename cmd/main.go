package main

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/calvinbui/homer-docker-service-discovery/internal/config"
	"github.com/calvinbui/homer-docker-service-discovery/internal/docker"
	"github.com/calvinbui/homer-docker-service-discovery/internal/homer"
	"github.com/calvinbui/homer-docker-service-discovery/internal/logger"
)

// https://github.com/traefik/traefik/blob/master/pkg/provider/docker/docker.go#L191
func main() {
	ctx := context.Background()

	logger.Init()
	logger.Info("Loading config")
	conf, err := config.New()
	if err != nil {
		logger.Fatal("Error parsing config", err)
	}

	logger.Debug("Retrieving Docker server version")
	serverVersion, err := docker.ServerVersion(conf.Docker)
	if err != nil {
		logger.Fatal("Failed to retrieve information of the Docker client and server host", err)
	}
	logger.Debug(fmt.Sprintf("Provider connection established with Docker %s (API %s)", serverVersion.Version, serverVersion.APIVersion))

	if c, err := homer.ReadConfig(*conf.HomerBaseConfig); err == nil {
		logger.Debug(fmt.Sprintf("Loaded Homer config from %s:\n%s", *conf.HomerBaseConfigPath, c))
	} else {
		logger.Fatal(fmt.Sprintf("Error reading config file %s", *conf.HomerBaseConfigPath), err)
	}

	logger.Info("Building Homer config")
	createConfig(ctx, *conf)

	logger.Info("Watching for Docker containers changes")
	eventsc, errc := docker.WatchEvents(ctx, conf.Docker)
	for {
		select {
		case event := <-eventsc:
			if event.Action == "create" || event.Action == "destroy" {
				logger.Info("A" + event.Action + " event occurred")
				logger.Info("Building Homer config")
				createConfig(ctx, *conf)
			}
		case err := <-errc:
			if errors.Is(err, io.EOF) {
				logger.Debug("Provider event stream closed")
			}
		case <-ctx.Done():
			return
		}
	}
}

func createConfig(ctx context.Context, conf config.Config) error {
	logger.Info("Getting Docker containers")
	containers, err := docker.ListRunningContainers(nil, conf.Docker)
	if err != nil {
		logger.Fatal("Failed to list containers for Docker", err)
	}

	var parsedContainers []docker.Container
	for _, container := range containers {
		parsedContainer, err := docker.ParseContainer(ctx, conf.Docker, container)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to inspect container %s", container.Names), err)
		}
		logger.Debug(fmt.Sprintf("Inspected container %s with labels %s", parsedContainer.Name, parsedContainer.Labels))
		parsedContainers = append(parsedContainers, parsedContainer)
	}

	logger.Info("Building Homer config")
	c, err := homer.BuildConfig(*conf.HomerBaseConfig, parsedContainers)
	if err != nil {
		logger.Fatal("Error building Homer config", err)
	}

	logger.Info("Updating Homer config")
	err = homer.PutConfig(c, *conf.HomerConfigPath, "777")
	if err != nil {
		logger.Fatal("Error updating Homer config file", err)
	}

	return nil
}
