package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	Otel    string      `json:"otel"`
	MongoDB MongoConfig `json:"mongodb"`
	Redis   struct {
		MovieReview   string `json:"moviereview"`
		UserReview    string `json:"userreview"`
		CastInfo      string `json:"castinfo"`
		MovieInfo     string `json:"movieinfo"`
		Plot          string `json:"plot"`
		ReviewStorage string `json:"reviewstorage"`
		User          string `json:"user"`
	} `json:"redis"`
	Service struct {
		CastInfo      string `json:"castinfo"`
		MovieInfo     string `json:"movieinfo"`
		Plot          string `json:"plot"`
		ReviewStorage string `json:"reviewstorage"`
		MovieReview   string `json:"moviereview"`
		UserReview    string `json:"userreview"`
		Rating        string `json:"rating"`
		ComposeReview string `json:"composereview"`
		User          string `json:"user"`
		Page          string `json:"page"`
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
