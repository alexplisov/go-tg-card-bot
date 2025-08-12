package config

import (
	"log/slog"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN,notEmpty,required"`
	OwnerChatID      int64  `env:"OWNER_CHAT_ID,notEmpty,required"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := godotenv.Load(); err != nil {
		slog.Warn("Con't load .env file", slog.Any("warn", err))
	}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
