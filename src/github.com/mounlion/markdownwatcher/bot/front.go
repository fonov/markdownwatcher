package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"fmt"
)

const (
	start = "Подписка оформлена"
)

func main() {

	fmt.Println("*****")

	bot, err := tgbotapi.NewBotAPI("***REMOVED***")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Text {
		case "/start":

			// Same logic db

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, start)

			bot.Send(msg)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}