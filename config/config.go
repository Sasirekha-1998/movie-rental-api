package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string `json:"port"`
}

// Config is the root config structure
type Config struct {
	Server ServerConfig `json:"server"`
}

// LoadConfig reads the configuration from the default JSON file path
func LoadConfig() (*Config, error) {
	const defaultPath = "config/config.json" // 👈 default file path

	data, err := os.ReadFile(defaultPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config JSON: %w", err)
	}

	return &cfg, nil
}
