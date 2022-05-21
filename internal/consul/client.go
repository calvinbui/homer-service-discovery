package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func CreateClient() (*api.Client, error) {
	fmt.Sprint("%w",api.DefaultConfig())
	client, err := api.NewClient(api.DefaultConfig())

	if err != nil {
		return nil, err
	}

	return client, nil
}
