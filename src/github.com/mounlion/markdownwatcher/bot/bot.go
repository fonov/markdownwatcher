package bot

import (
	"github.com/mounlion/markdownwatcher/parsing"
	"fmt"
	"github.com/mounlion/markdownwatcher/database"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"github.com/mounlion/markdownwatcher/model"
)

const DNSDomain = "https://www.dns-shop.ru"
var BotToken *string

func SetBotToken(value *string)  {
	BotToken = value
}

func SendCatalog(newItems []parsing.Item, updateItems []model.UpdateItem)  {
	bot, err := tgbotapi.NewBotAPI(*BotToken)
	if err != nil {
		log.Panic(err)
	}

	users := database.GetUsers()

	if len(newItems) > 0 || len(updateItems) > 0 {

		var (
			newItemsString string
			updateItemsString string
		)

		if len(newItems) > 0 {
			newItemsString += "<b>Новые товары</b>\n\n"
			for _, val := range newItems {
				newItemsString += CatalogMessage(val, 0)
			}
		}

		if len(updateItems) > 0 {
			updateItemsString += "<b>Обновление цен</b>\n\n"
			for _, val := range updateItems {
				updateItemsString += CatalogMessage(val.Item, val.OldDiDiscountPrice)
			}
		}

		for _, user := range users {
			if user.IsActive {
				if len(newItemsString) > 0 {
					sendMessage(bot, &user, &newItemsString, false)
				}
				if len(updateItemsString) > 0 {
					sendMessage(bot, &user, &updateItemsString, false)
				}
			}
		}
	} else {
		for _, user := range users {
			if user.IsActive {
				text := "<b>Новых или обновленных товаров не найдено</b>"
				sendMessage(bot, &user, &text, true)
			}
		}
	}
}

func sendMessage(bot *tgbotapi.BotAPI, user *database.User, message *string, DisableNotification bool)  {
	msg := tgbotapi.NewMessage(user.Id, *message)
	msg.ParseMode = "HTML"
	msg.DisableWebPagePreview = true
	msg.DisableNotification = DisableNotification
	_, err := bot.Send(msg)
	if err != nil {
		if err.Error() == "Forbidden: bot was blocked by the user" {
			database.Subscribe(int(user.Id), false)
		}
	}
}

func SendServiceMessage(text string)  {
	users := database.GetUsers()

	bot, err := tgbotapi.NewBotAPI(*BotToken)
	if err != nil {
		log.Panic(err)
	}

	for _, user := range users {
		if user.IsActive && user.IsAdmin {
			sendMessage(bot, &user, &text, false)
		}
	}
}

func CatalogMessage(item parsing.Item, OldDiDiscountPrice int)string {
	var catalog string

	catalog += fmt.Sprintf("<a href=\"%s%s\">%s</a>\n", DNSDomain, item.Url, item.Title)
	catalog += fmt.Sprintf("<b>%d₽</b>", item.Price)
	if item.OldPrice != 0 {
		profit := 100-(float64(item.Price)/float64(item.OldPrice)*100)
		catalog += fmt.Sprintf("    <code>%d₽ %.1f%%</code>", item.OldPrice, profit)
	}
	if OldDiDiscountPrice != 0 {
		catalog += fmt.Sprintf("\n<i>Переуценка на %d₽</i>", OldDiDiscountPrice-item.Price)
	}
	if len(item.Desc) > 0 {
		catalog += fmt.Sprintf("<i>%s</i>", item.Desc)
	}
	catalog += "\n\n"

	return catalog
}