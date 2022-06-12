package consul

import (
	"github.com/hashicorp/consul/api"
)

func CreateClient(host string) (*api.Client, error) {
	client, err := api.NewClient(&api.Config{Address: host})

	if err != nil {
		return nil, err
	}

	return client, nil
}
