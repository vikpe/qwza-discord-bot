package config

import (
	"github.com/goccy/go-json"
	"os"
)

type Config struct {
	Monitor Monitor `json:"monitor"`
}

type Monitor struct {
	Interval int      `json:"interval"`
	Servers  []string `json:"servers"`
	Channel  string   `json:"channel"`
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
