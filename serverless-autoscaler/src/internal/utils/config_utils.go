package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	Interval  int      `json:"interval"` // check interval in seconds
	Namespace string   `json:"namespace"`
	Apps      []string `json:"apps"`
}

func ParseConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
