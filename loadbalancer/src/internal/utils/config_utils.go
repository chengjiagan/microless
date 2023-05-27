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
	Enable        bool `json:"enable"`
	Reject        bool `json:"reject"`
	MaxTokens     int  `json:"max_tokens"`
	TokensPerFill int  `json:"tokens_per_fill"`
	FillInterval  int  `json:"fill_interval"`
}

type ClientConfig struct {
	Enable            bool              `json:"enable"`
	VmPostfix         string            `json:"vm_postfix"`
	ServerlessPostfix string            `json:"serverless_postfix"`
	DegradeInterval   int               `json:"degrade_interval"`
	LocalServices     map[string]string `json:"local_services"`
	Retry             int               `json:"retry"`
}

type ServerlessConfig struct {
	Enable            bool           `json:"enable"`
	MaxConcurrency    int            `json:"max_concurrency"`
	MaxCapacity       int            `json:"max_capacity"`
	MethodReqirements map[string]int `json:"method_requirements"` // if method requires a full CPU core, then its requirement is 100
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
