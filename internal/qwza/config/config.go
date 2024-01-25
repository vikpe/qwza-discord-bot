package config

import (
	"github.com/goccy/go-json"
	"os"
)

type Config struct {
	Monitor Monitor `json:"monitor"`
}

type Monitor struct {
	Interval  uint64   `json:"interval"`
	Servers   []string `json:"servers"`
	ChannelId string   `json:"channel_id"`
}

func FromFile(path string) (*Config, error) {
	bytes, err := os.ReadFile(path)

	if err != nil {
		return &Config{}, err
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	return &config, nil
}
