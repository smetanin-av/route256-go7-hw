package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	pathToConfig = "config.yaml"
)

type Config struct {
	GrpcPort       uint32 `yaml:"grpc_port"`
	HttpPort       uint32 `yaml:"http_port"`
	Token          string `yaml:"token"`
	LomsService    string `yaml:"loms_service"`
	ProductService string `yaml:"product_service"`
	DatabaseDSN    string `yaml:"database_dsn"`
	RedisAddr      string `yaml:"redis_addr"`
}

func NewConfig() (*Config, error) {
	data, err := os.ReadFile(pathToConfig)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &config, nil
}
