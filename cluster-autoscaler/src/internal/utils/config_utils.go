package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	Interval       int           `json:"interval"`        // check interval in seconds
	StableInterval int           `json:"stable_interval"` // stable interval in minutes
	Namespace      string        `json:"namespace"`
	Ratio          int           `json:"ratio"` // serverless/vm ratio for same rps
	Kube           KubeConfig    `json:"kube"`
	Cluster        ClusterConfig `json:"cluster"`
}

type KubeConfig struct {
	VmSelector         string `json:"vm_selector"`
	ServerlessSelector string `json:"serverless_selector"`
}

type ClusterConfig struct {
	ServerlessCpu float64 `json:"serverless_cpu"`
	ServerlessMem float64 `json:"serverless_mem"`
	VmCpu         float64 `json:"vm_cpu"`
	VmMem         float64 `json:"vm_mem"`

	ServerlessCpuPrice float64 `json:"serverless_cpu_price"`
	ServerlessMemPrice float64 `json:"serverless_mem_price"`
	VmPrice            float64 `json:"vm_price"`
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
