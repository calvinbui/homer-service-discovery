package docker

func (c Container) GetLabelValue(label string) (string, bool) {
	value, ok := c.Labels[label]
	return value, ok
}

func (c Container) GetLabelValueOrEmpty(label string) string {
	if value, ok := c.Labels[label]; ok {
		return value
	}

	return ""
}
