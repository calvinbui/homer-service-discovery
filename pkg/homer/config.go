package homer

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/calvinbui/homer-docker-service-discovery/internal/docker"
	"github.com/calvinbui/homer-docker-service-discovery/internal/logger"
)

const (
	EnableLabel     = "homer.enable"
	ServiceLabel    = "homer.service"
	NameLabel       = "homer.name"
	LogoLabel       = "homer.logo"
	IconLabel       = "homer.icon"
	SubtitleLabel   = "homer.subtitle"
	TagLabel        = "homer.tag"
	UrlLabel        = "homer.url"
	TargetLabel     = "homer.target"
	TagstyleLabel   = "homer.tagstyle"
	TypeLabel       = "homer.type"
	ClassLabel      = "homer.class"
	BackgroundLabel = "homer.background"
	PriorityLabel   = "homer.priority"
)

func BuildConfig(config Config, containers []docker.Container) (Config, error) {
	logger.Debug("Checking all container and their labels")
	for _, container := range containers {
		logger.Debug(fmt.Sprintf("Start checking Container %s", container.Name))

		logger.Debug(fmt.Sprintf("Container %s: Checking if label %s exists and is true", container.Name, EnableLabel))
		if homerEnabled(container) {
			logger.Debug(fmt.Sprintf("Container %s: Label %s exists and is true", container.Name, EnableLabel))

			logger.Debug(fmt.Sprintf("Container %s: Finding Homer Service from label %s", container.Name, ServiceLabel))

			if i, found := findServiceFromLabel(container, config.Services); found {
				logger.Debug(fmt.Sprintf("Container %s: Found matching Homer Service %s", container.Name, config.Services[i].Name))

				logger.Debug(fmt.Sprintf("Container %s: Appending to Homer Service %s", container.Name, config.Services[i].Name))
				config.Services[i].Items = append(config.Services[i].Items, Item{
					Name:       container.GetLabelValueOrEmpty(NameLabel),
					Logo:       container.GetLabelValueOrEmpty(LogoLabel),
					Icon:       container.GetLabelValueOrEmpty(IconLabel),
					Subtitle:   container.GetLabelValueOrEmpty(SubtitleLabel),
					Tag:        container.GetLabelValueOrEmpty(TagLabel),
					Url:        container.GetLabelValueOrEmpty(UrlLabel),
					Target:     container.GetLabelValueOrEmpty(TargetLabel),
					Tagstyle:   container.GetLabelValueOrEmpty(TagstyleLabel),
					Type:       container.GetLabelValueOrEmpty(TypeLabel),
					Class:      container.GetLabelValueOrEmpty(ClassLabel),
					Background: container.GetLabelValueOrEmpty(BackgroundLabel),
					Priority:   toIntOrZero(container.GetLabelValueOrEmpty(PriorityLabel)),
				})
			} else {
				logger.Warn(fmt.Sprintf("Container %s: No matching service found or label does not exist. Skipping", container.Name), nil)
				continue
			}
		} else {
			logger.Debug(fmt.Sprintf("Container %s: Label %s does not exist or is false", container.Name, EnableLabel))
		}
	}

	logger.Debug("Sorting all services based on priority")
	for i := range config.Services {
		logger.Debug(fmt.Sprintf("Service %s: Sorting items", config.Services[i].Name))

		sort.Slice(config.Services[i].Items, func(j, k int) bool {
			return config.Services[i].Items[j].Priority > config.Services[i].Items[k].Priority
		})
	}

	return config, nil
}

// homerEnabled checks if the container has the enabled label and it is true
func homerEnabled(c docker.Container) bool {
	value, ok := c.Labels[EnableLabel]
	return ok && value == "true"
}

// findServiceFromLabel finds the service specified on the container's labels and returns its position
func findServiceFromLabel(c docker.Container, services []Service) (int, bool) {
	if value, found := c.GetLabelValue(ServiceLabel); found {
		for index, service := range services {
			if service.Name == value {
				return index, true
			}
		}
	}

	return 0, false
}

func toIntOrZero(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return i
}
