package entry

type RawEntry struct {
	Name   string
	Labels map[string]string
}

func (c RawEntry) GetLabelValue(label string) (string, bool) {
	value, ok := c.Labels[label]
	return value, ok
}

func (c RawEntry) GetLabelValueOrEmpty(label string) string {
	if value, ok := c.Labels[label]; ok {
		return value
	}

	return ""
}
