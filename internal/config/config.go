package config

import (
	yaml "gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	StockHostName       string        `yaml:"STOCK_API_HOST_NAME"`
	StockHostPort       string        `yaml:"STOCK_API_HOST_PORT"`
	StockAPIConnTimeout time.Duration `yaml:"stop_api_connection_timeout"`
}

const (
	configFilePathEnv  = "CONFIG_FILE"
	defaultFilePathEnv = "config.yaml"
)

func Load() (Config, error) {
	var cfg Config

	configFilePath, exists := os.LookupEnv(configFilePathEnv)
	if !exists {
		configFilePath = defaultFilePathEnv
	}

	filename, _ := filepath.Abs(configFilePath)
	yamlFile, err := os.ReadFile(filename)

	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
