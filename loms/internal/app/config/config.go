package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	pathToConfig = "config.yaml"
)

var (
	ErrMaxAwaitingPaymentTooSmall = errors.New("max_awaiting_payment min value is 1s")
)

type Config struct {
	GrpcPort           uint32        `yaml:"grpc_port"`
	HttpPort           uint32        `yaml:"http_port"`
	DatabaseDSN        string        `yaml:"database_dsn"`
	MaxAwaitingPayment time.Duration `yaml:"max_awaiting_payment"`
	KafkaBrokers       []string      `yaml:"kafka_brokers"`
	KafkaTopic         string        `yaml:"kafka_topic"`
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

	if config.MaxAwaitingPayment < time.Second {
		return nil, ErrMaxAwaitingPaymentTooSmall
	}

	return &config, nil
}
