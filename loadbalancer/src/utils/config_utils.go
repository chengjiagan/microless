package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	Grpc    string `json:"grpc"`
	Http    string `json:"http"`
	Backend struct {
		Vm         string `json:"vm"`
		Serverless string `json:"serverless"`
	} `json:"backend"`
	Deployment *DeploymentConfig `json:"deployment"`
	RatePerPod int               `json:"rate_per_pod"`
}

type DeploymentConfig struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
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
