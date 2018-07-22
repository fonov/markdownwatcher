package bot

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"github.com/mounlion/markdownwatcher/database"
)

const (
	botToken = "***REMOVED***"
	start = "Подписка оформлена"
	stop = "Подписка отмена"
	otherwise = "Введена неправильная команда"
)

func Front() {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

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
			bot.Send(msg)
		case "/stop":
			database.Subscribe(update.Message.From.ID, false)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, stop)
			bot.Send(msg)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, otherwise)
			bot.Send(msg)
		}
	}
}