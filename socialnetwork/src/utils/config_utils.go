package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	Grpc    string      `json:"grpc"`
	Rest    string      `json:"rest"`
	Otel    string      `json:"otel"`
	MongoDB MongoConfig `json:"mongodb"`
	Redis   struct {
		UserTimeline string `json:"usertimeline"`
		SocialGraph  string `json:"socialgraph"`
		HomeTimeline string `json:"hometimeline"`
		PostStorage  string `json:"poststorage"`
		User         string `json:"user"`
		UrlShorten   string `json:"urlshorten"`
	} `json:"redis"`
	Service struct {
		PostStorage  string `json:"poststorage"`
		UserTimeline string `json:"usertimeline"`
		User         string `json:"user"`
		SocialGraph  string `json:"socialgraph"`
		HomeTimeline string `json:"hometimeline"`
		Media        string `json:"media"`
		UrlShorten   string `json:"urlshorten"`
		UserMention  string `json:"usermention"`
		Text         string `json:"text"`
		ComposePost  string `json:"composepost"`
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
