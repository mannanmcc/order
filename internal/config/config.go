package config

import (
	"errors"
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
	environment        = "ENVIRONMENT"
)

func Load() (Config, error) {
	var cfg Config

	if os.Getenv(environment) == "production" {
		return readFromEnvironment(cfg)
	}
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

func readFromEnvironment(cfg Config) (Config, error) {
	cfg.StockHostName = os.Getenv("STOCK_API_HOST_NAME")
	cfg.StockHostPort = os.Getenv("STOCK_API_HOST_PORT")
	if cfg.StockHostName == "" {
		return cfg, errors.New("STOCK_API_HOST_NAME environment variable not set")
	}
	if cfg.StockHostPort == "" {
		return cfg, errors.New("STOCK_API_HOST_PORT environment variable not set")
	}

	return cfg, nil
}
