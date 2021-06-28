package docker

import (
	"github.com/docker/docker/client"
)

func CreateClient() (*client.Client, error) {
	client, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return nil, err
	}

	return client, nil
}
