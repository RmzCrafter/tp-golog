package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func LoadConfig(path string) ([]LogConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var configs []LogConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		return nil, err
	}

	return configs, nil
}

func SaveConfig(configs []LogConfig, path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(configs, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func AddLogToConfig(configPath string, newLog LogConfig) error {
	configs, err := LoadConfig(configPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			configs = []LogConfig{}
		} else {
			return err
		}
	}

	for _, config := range configs {
		if config.ID == newLog.ID {
			return fmt.Errorf("un log avec l'ID '%s' existe déjà", newLog.ID)
		}
	}

	configs = append(configs, newLog)
	return SaveConfig(configs, configPath)
} 