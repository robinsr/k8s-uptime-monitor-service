// Loads config from file/env

package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Target struct {
	Name     string `yaml:"name"`     // Friendly name for the service
	URL      string `yaml:"url"`      // URL to monitor
	Interval int    `yaml:"interval"` // Ping interval in seconds
	Timeout  int    `yaml:"timeout"`  // Timeout per request in seconds
}

type Config struct {
	Targets []Target `yaml:"targets"`
}

func LoadConfig(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
