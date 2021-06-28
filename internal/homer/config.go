package homer

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func readConfig(path string) ([]byte, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func unmarshalConfig(contents []byte) (Config, error) {
	config := Config{}

	if len(contents) == 0 {
		return Config{}, fmt.Errorf("Homer config is empty")
	}

	err := yaml.Unmarshal(contents, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
