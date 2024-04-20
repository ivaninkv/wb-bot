package main

import (
	"wb-bot/bot"
	"wb-bot/config"
	"wb-bot/logger"
)

func main() {
	logger.Log.Info("Starting app...")
	cfg := config.LoadConfig()

	bot.Start(cfg)
}
