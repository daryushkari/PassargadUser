package config

import (
	"encoding/json"
	"os"
)

type SecretData struct {
	JWTSecret string `json:"jwt-secret"`
}

var SampleSecretKey []byte

func SetSecret(cnfPath string) error {
	secretCFG := &SecretData{}
	file, err := os.ReadFile(cnfPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, secretCFG)
	if err != nil {
		return err
	}
	SampleSecretKey = []byte(secretCFG.JWTSecret)
	return nil
}
