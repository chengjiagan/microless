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
		Bookings string `json:"bookings"`
		Customer string `json:"customer"`
		Flights  string `json:"flights"`
	} `json:"memcached"`
	Service struct {
		Bookings string `json:"bookings"`
		Customer string `json:"customer"`
		Flights  string `json:"flights"`
	} `json:"service"`
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
