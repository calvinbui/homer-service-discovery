package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Container struct {
	ID     string
	Name   string
	Labels map[string]string
}

func ListRunningContainers(ctx context.Context, docker client.APIClient) ([]types.Container, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	containers, err := docker.ContainerList(ctx, types.ContainerListOptions{All: false})

	if err != nil {
		return nil, err
	}

	return containers, nil
}

func ParseContainer(ctx context.Context, docker client.APIClient, container types.Container) (Container, error) {
	i, err := docker.ContainerInspect(ctx, container.ID)

	if err != nil {
		return Container{}, err
	}

	c := Container{
		ID:   i.ContainerJSONBase.ID,
		Name: i.ContainerJSONBase.Name,
	}

	if i.Config != nil && i.Config.Labels != nil {
		c.Labels = i.Config.Labels
	}

	return c, nil
}
