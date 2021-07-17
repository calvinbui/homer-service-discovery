package homer

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"

	"gopkg.in/yaml.v3"
)

func GetConfig(path string) (*Config, error) {
	b, err := openConfig(path)
	if err != nil {
		return nil, err
	}

	config, err := unmarshalConfig(b)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling config file: %w", err)
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

	// the footer can be disabled by changing it from a 'string' to a 'bool'
	if config.Footer == "false" {
		config.Footer = false
	}

	return config, nil
}

func ReadConfig(config Config) ([]byte, error) {
	b, err := yaml.Marshal(config)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func PutConfig(config Config, path string, permissions string) error {
	b, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("Error marshalling data into YAML: %w", err)
	}

	fileStats, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("Error getting stats for config file '%s': %w", path, err)
	}

	err = ioutil.WriteFile(path, b, fileStats.Mode())
	if err != nil {
		return fmt.Errorf("Error writing data to config file '%s'", path)
	}

	fileSysStats := fileStats.Sys().(*syscall.Stat_t)
	uid := int(fileSysStats.Uid)
	gid := int(fileSysStats.Gid)
	err = os.Chown(path, uid, gid)
	if err != nil {
		return fmt.Errorf("Error chowning config '%s' with user '%v' and group '%v': %w", path, uid, gid, err)
	}

	return nil
}
