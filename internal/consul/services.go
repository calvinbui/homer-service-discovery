package consul

import (
	"github.com/calvinbui/homer-docker-service-discovery/internal/entry"
	"github.com/hashicorp/consul/api"
	"strings"
)

func ListServices(consul *api.Client) map[string][]string {

	catalog := consul.Catalog()
	q := &api.QueryOptions{}
	services, _, _ := catalog.Services(q)
	return services
}

func ParseService(name string, labels []string) entry.RawEntry {
	s := entry.RawEntry{
		Name: name,
	}
	s.Labels = make(map[string]string)
	if labels != nil && len(labels) > 0 {
		for _, label := range labels {
			splitedLabel := strings.Split(label, "=")
			if len(splitedLabel) > 1 {
				s.Labels[splitedLabel[0]] = splitedLabel[1]
			} else {
				s.Labels[splitedLabel[0]] = ""
			}

		}
	}
	return s
}
