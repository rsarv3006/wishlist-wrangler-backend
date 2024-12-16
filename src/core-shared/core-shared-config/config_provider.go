package coresharedconfig

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

func LoadConfig(serviceName string) (ConfigDto, error) {
	_ = serviceName
	return loadFromLocalFile()
}

func loadFromLocalFile() (ConfigDto, error) {
	var config ConfigDto

	data, err := os.ReadFile("./config.yml")
	if err != nil {
		return config, fmt.Errorf("error reading config file: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error parsing config file: %v", err)
	}

	return config, nil
}
