package docker

import (
	"context"

	"github.com/calvinbui/homer-docker-service-discovery/internal/entry"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

func ListRunningServices(ctx context.Context, docker client.APIClient) ([]swarm.Service, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	services, err := docker.ServiceList(ctx, types.ServiceListOptions{Status: true})

	if err != nil {
		return nil, err
	}

	return services, nil
}

func ParseService(ctx context.Context, docker client.APIClient, service swarm.Service) (entry.RawEntry, error) {
	i, _, err := docker.ServiceInspectWithRaw(ctx, service.ID, types.ServiceInspectOptions{})

	if err != nil {
		return entry.RawEntry{}, err
	}

	s := entry.RawEntry{
		Name: i.Spec.Annotations.Name,
	}

	if i.Spec.Labels != nil {
		s.Labels = i.Spec.Annotations.Labels
	}

	return s, nil
}
