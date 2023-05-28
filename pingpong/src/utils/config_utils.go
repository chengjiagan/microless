package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	Service struct {
		Ping string `json:"ping"`
		Pong string `json:"pong"`
	} `json:"service"`
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
