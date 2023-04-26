package config

import (
	"encoding/json"
	"os"
)

var (
	config *EnvConfig
)

type EnvConfig struct {
	Database struct {
		Name string `json:"name"`
	} `json:"database"`
	ExternalExpose struct {
		GrpcPort string `json:"grpc-port"`
		RestPort string `json:"rest-port"`
	} `json:"externalExpose"`
}

type FullConfig struct {
	Prod EnvConfig `json:"production"`
	Test EnvConfig `json:"test"`
}

const (
	ProductionEnv = "production"
	TestEnv       = "test"
)

func Get(cnfPath string, env string) (*EnvConfig, error) {
	if config == nil {
		config = &EnvConfig{}
		fullCfg := map[string]*EnvConfig{}
		file, err := os.ReadFile(cnfPath)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(file, &fullCfg)
		if err != nil {
			return nil, err
		}
		config = fullCfg[env]
	}
	return config, nil
}
