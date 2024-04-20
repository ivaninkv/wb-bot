package config

import (
	"os"
	"wb-bot/logger"
)

type Config struct {
	BotToken string `json:"bot_token"`
}

func LoadConfig() *Config {
	logger.Log.Info("Loading config...")
	config := &Config{
		BotToken: os.Getenv("BOT_TOKEN"),
	}

	if config.BotToken == "" {
		panic("BOT_TOKEN is not set")
	}

	logger.Log.Info("Config loaded")

	return config
}
