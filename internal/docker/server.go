package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ServerVersion(c *client.Client) (types.Version, error) {
	return c.ServerVersion(context.TODO())
}
