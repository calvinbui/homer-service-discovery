package consul

import (
	"github.com/calvinbui/homer-docker-service-discovery/internal/logger"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
)

func WatchServices(consul *api.Client) *watch.Plan {
	servicesWatcher, err := watch.Parse(map[string]interface{}{"type": "services"})
	if err != nil {
		logger.Fatal("failed to create services watcher plan: %w", err)
	}
	return servicesWatcher
}
