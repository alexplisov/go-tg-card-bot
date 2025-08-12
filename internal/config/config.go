package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
