package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	Grpc      string      `json:"grpc"`
	Rest      string      `json:"rest"`
	Otel      string      `json:"otel"`
	MongoDB   MongoConfig `json:"mongodb"`
	Memcached struct {
		Profile     string `json:"profile"`
		Rate        string `json:"rate"`
		User        string `json:"user"`
		Reservation string `json:"reservation"`
	} `json:"memcached"`
	Redis struct {
		Geo string `json:"geo"`
	} `json:"redis"`
	Service struct {
		Geo         string `json:"geo"`
		Profile     string `json:"profile"`
		Rate        string `json:"rate"`
		Reservation string `json:"reservation"`
		Search      string `json:"search"`
		User        string `json:"user"`
	} `json:"service"`
	Secret string `json:"secret"`
}

type MongoConfig struct {
	Url      string `json:"url"`
	Database string `json:"database"`
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
