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
	DatabaseDSN   string   `yaml:"database_dsn"`
	KafkaBrokers  []string `yaml:"kafka_brokers"`
	KafkaTopic    string   `yaml:"kafka_topic"`
	ConsumerGroup string   `yaml:"consumer_group"`
	TgBotApiToken string   `yaml:"tg_bot_api_token"`
	TgBotChatId   int64    `yaml:"tg_bot_chat_id"`
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
