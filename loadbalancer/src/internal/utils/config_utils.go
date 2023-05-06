package utils

import (
	"encoding/json"
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
	MaxTokens     int  `json:"max_tokens"`
	TokensPerFill int  `json:"tokens_per_fill"`
	FillInterval  int  `json:"fill_interval"`
}

type ClientConfig struct {
	Enable            bool   `json:"enable"`
	VmPostfix         string `json:"vm_postfix"`
	ServerlessPostfix string `json:"serverless_postfix"`
	DegradeInterval   int    `json:"degrade_interval"`
}

type ServerlessConfig struct {
	Enable bool `json:"enable"`
}

var configPath = os.Getenv("LB_CONFIG")
var config *Config

func getConfig() *Config {
	if config != nil {
		return config
	}

	data, err := os.ReadFile(configPath)
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
	config := getConfig()
	return &config.Server
}

func GetClientConfig() *ClientConfig {
	config := getConfig()
	return &config.Client
}

func GetServerlessConfig() *ServerlessConfig {
	config := getConfig()
	return &config.Serverless
}
