package homer

import (
	"fmt"
	"io/ioutil"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

func GetConfig(path string) (*Config, error) {
	b, err := openConfig(path)
	if err != nil {
		return nil, err
	}

	config, err := unmarshalConfig(b)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func openConfig(path string) ([]byte, error) {
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

func ReadConfig(config Config) ([]byte, error) {
	b, err := yaml.Marshal(config)
	if err != nil {
		return nil, err
	}

	log.Print("Hi")

	return b, nil
}
