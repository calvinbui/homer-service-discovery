package main

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/calvinbui/homer-docker-service-discovery/internal/config"
	"github.com/calvinbui/homer-docker-service-discovery/internal/docker"
	"github.com/calvinbui/homer-docker-service-discovery/internal/logger"
	"github.com/calvinbui/homer-docker-service-discovery/pkg/homer"
)

func main() {
	logger.Init()

	logger.Debug("Loading internal config")
	conf, err := config.New()
	if err != nil {
		logger.Fatal("Error parsing config", err)
	}

	logger.Debug("Retrieving Docker server version")
	serverVersion, err := docker.ServerVersion(conf.Docker)
	if err != nil {
		logger.Fatal("Failed to retrieve information about the Docker client and server host", err)
	}
	logger.Trace(fmt.Sprintf("Provider connection established with Docker %s (API %s)", serverVersion.Version, serverVersion.APIVersion))

	logger.Info("Building Homer config from base config")
	ctx := context.Background()
	err = generateConfig(ctx, conf)
	if err != nil {
		logger.Fatal("Error generating Homer config", err)
	}

	logger.Info("Start watching for container creations and deletions")
	eventsc, errc := docker.WatchEvents(ctx, conf.Docker)
	for {
		select {
		case event := <-eventsc:
			if event.Action == "create" || event.Action == "destroy" {
				logger.Trace(fmt.Sprintf("%+v", event))
				logger.Debug("A " + event.Action + " event occurred")
				logger.Info(fmt.Sprintf("Change detected. %s was %s, generating Homer config...", event.From, event.Action))
				err = generateConfig(ctx, conf)
				if err != nil {
					logger.Fatal("Error generating Homer config", err)
				}
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

func generateConfig(ctx context.Context, conf config.Config) error {
	logger.Debug("Getting Docker containers")
	containers, err := docker.ListRunningContainers(ctx, conf.Docker)
	if err != nil {
		logger.Fatal("Failed to list containers for Docker", err)
	}
	var parsedContainers []docker.Container
	for _, container := range containers {
		parsedContainer, err := docker.ParseContainer(ctx, conf.Docker, container)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to inspect container %s", container.Names), err)
		}
		logger.Debug(fmt.Sprintf("Inspected container %s", parsedContainer.Name))
		parsedContainers = append(parsedContainers, parsedContainer)
	}

	logger.Debug("Generating config")
	baseConfig, err := conf.HomerBaseConfig.DeepCopy()
	if err != nil {
		logger.Fatal("Error creating a deep copy of the base config", err)
	}
	generatedConfig, err := homer.BuildConfig(baseConfig, parsedContainers)
	if err != nil {
		logger.Fatal("Error building Homer config", err)
	}

	logger.Debug("Updating Homer config file")
	err = homer.PutConfig(generatedConfig, conf.HomerConfigPath, "777")
	if err != nil {
		logger.Fatal("Error updating Homer config file", err)
	}

	logger.Info("Homer config was successfully generated and updated")

	return nil
}
