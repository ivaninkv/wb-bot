package bot

import (
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
	"strings"
	"time"
	"wb-bot/config"
	"wb-bot/logger"
	"wb-bot/wb"
)

func Start(cfg *config.Config) {
	b, err := telebot.NewBot(telebot.Settings{
		Token:  cfg.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		logger.Log.Error(err.Error())
		panic(err)
	}

	b.Handle("/start", StartHandler(b))
	b.Handle("/help", StartHandler(b))
	b.Handle(telebot.OnText, SearchHandler(b))

	b.Start()
}

func StartHandler(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		logger.Log.Info(fmt.Sprintf("Received %v command from user %s", m.Text, m.Sender.Username))
		_, err := b.Send(m.Sender, "Для поиска артикулов на WB напишите поисковый запрос в чат")
		if err != nil {
			logger.Log.Error(err.Error())
			return
		}
	}
}

func SearchHandler(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		logger.Log.Info(fmt.Sprintf("Received search query '%s' from user %s", m.Text, m.Sender.Username))
		response, err := wb.GetProductIDs(m.Text)
		if err != nil {
			logger.Log.Error(fmt.Sprintf("Error occurred while searching for product IDs: %v", err))
			_, err := b.Send(m.Sender, "Что-то пошло не так")
			if err != nil {
				logger.Log.Error(err.Error())
				return
			}
			return
		}
		stringSlice := make([]string, len(response))
		for i, num := range response {
			stringSlice[i] = strconv.Itoa(num)
		}
		logger.Log.Info(fmt.Sprintf("Found %d product IDs for query '%s'", len(response), m.Text))
		_, err = b.Send(m.Sender, strings.Join(stringSlice, "\n"))
		if err != nil {
			logger.Log.Error(err.Error())
			return
		}
	}
}
