package config

import (
	"encoding/json"
	"os"
)

var (
	config *Config = nil
)

type Config struct {
	Database struct {
		Name string `json:"name"`
	} `json:"database"`
	ExternalExpose struct {
		GrpcPort string `json:"grpc-port"`
		RestPort string `json:"rest-port"`
	} `json:"externalExpose"`
}

func Get(cnfPath string) (*Config, error) {
	if config == nil {
		file, err := os.ReadFile(cnfPath)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(file, config)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}
