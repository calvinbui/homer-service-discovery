package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func WatchEvents(ctx context.Context, docker client.APIClient, swarmMode bool) (<-chan events.Message, <-chan error) {
	f := filters.NewArgs()
	f.Add("type", "container")

	if swarmMode {
		f.Add("type", "service")
	}

	return docker.Events(ctx, types.EventsOptions{Filters: f})
}
