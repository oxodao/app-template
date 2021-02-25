package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {

	Web struct {
		ListeningHost string
	}

}

func LoadConfig() (*Config, error) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(file, &cfg)

	return &cfg, err
}