// Author: Vincent Composieux <vincent.composieux@gmail.com>

package app

import (
	"encoding/json"
	"fmt"
	"os"
)

// JWTConfig is the "jwt" configuration section type definition.
type JWTConfig struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

// Config is the configuration type definition.
type Config struct {
	JWT JWTConfig `json:"jwt"`
}

// GetConfig returns the configuration object that can be used anywhere in application.
func GetConfig() Config {
	file, _ := os.Open("app/config.json")
	decoder := json.NewDecoder(file)

	config := Config{}
	err := decoder.Decode(&config)

	if err != nil {
		fmt.Println("An error occurs on configuration loading:", err)
	}

	return config
}
