package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"homework-2/config"
	"homework-2/rss"
	"log"
	"os"
)

func main() {

	rawСfg, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.ParseConfig(rawСfg)
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.ApiKeys.Telegram)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, rss.Try(update.Message.Text))

		bot.Send(msg)
	}
}
