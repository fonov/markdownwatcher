package bot

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"github.com/mounlion/markdownwatcher/database"
	"github.com/mounlion/markdownwatcher/config"
)

const (
	start = "Подписка оформлена"
	stop = "Подписка отмена"
	otherwise = "Введена неправильная команда"
)

// Front end of telegram boot process users command
func Front() {
	bot, err := tgbotapi.NewBotAPI(*config.Config.BotToken)
	if err != nil {
		log.Panic(err)
	}

	if *config.Config.Logger {log.Printf("Start bot front. Token: %s", *config.Config.BotToken)}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Text {
		case "/start":
			database.Subscribe(update.Message.From.ID, true)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, start)
			_, err := bot.Send(msg)
			if err != nil {
				log.Print(err.Error())
			}
		case "/stop":
			database.Subscribe(update.Message.From.ID, false)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, stop)
			_, err := bot.Send(msg)
			if err != nil {
				log.Print(err.Error())
			}
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, otherwise)
			_, err := bot.Send(msg)
			if err != nil {
				log.Print(err.Error())
			}
		}
	}
}