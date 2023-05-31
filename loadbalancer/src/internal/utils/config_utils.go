package utils

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

type Config struct {
	Server     ServerConfig     `json:"server"`
	Client     ClientConfig     `json:"client"`
	Serverless ServerlessConfig `json:"serverless"`
}

type ServerConfig struct {
	Enable         bool    `json:"enable"`
	UpdateInterval int     `json:"update_interval"`
	UpdateRatio    float64 `json:"update_ratio"`
}

type ClientConfig struct {
	Enable            bool              `json:"enable"`
	VmPostfix         string            `json:"vm_postfix"`
	ServerlessPostfix string            `json:"serverless_postfix"`
	UpdateInterval    int               `json:"update_interval"`
	Retry             int               `json:"retry"`
	ServiceRateLimit  map[string]int    `json:"service_rate_limit"`
	LocalServices     map[string]string `json:"local_services"`
	RedisAddr         string            `json:"redis_addr"`
}

type ServerlessConfig struct {
	Enable            bool               `json:"enable"`
	MaxConcurrency    int                `json:"max_concurrency"`
	MaxCapacity       int                `json:"max_capacity"`
	MethodReqirements map[string]float64 `json:"method_requirements"`
}

var configPath = flag.String("lb_config", os.Getenv("LB_CONFIG"), "path to loadbalancer config file")
var config *Config

func GetConfig() *Config {
	if config != nil {
		return config
	}

	data, err := os.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	config = new(Config)
	err = json.Unmarshal(data, config)
	if err != nil {
		config = nil
		log.Fatalf("Failed to parse config file: %v", err)
	}

	return config
}

func GetServerConfig() *ServerConfig {
	config := GetConfig()
	return &config.Server
}

func GetClientConfig() *ClientConfig {
	config := GetConfig()
	return &config.Client
}

func GetServerlessConfig() *ServerlessConfig {
	config := GetConfig()
	return &config.Serverless
}
