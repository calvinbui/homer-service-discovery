package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func WatchEvents(ctx context.Context, docker client.APIClient) (<-chan events.Message, <-chan error) {
	f := filters.NewArgs()
	f.Add("type", "container")

	return docker.Events(ctx, types.EventsOptions{Filters: f})
}
