package bot

import (
	"github.com/mounlion/markdownwatcher/parsing"
	"fmt"
	"github.com/mounlion/markdownwatcher/database"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

const DNSDomain = "https://www.dns-shop.ru"

func SendMessage(newItems []parsing.Item, updateItems[]parsing.Item)  {
	if len(newItems) > 0 || len(updateItems) > 0 {

		var (
			newItemsString string
			updateItemsString string
		)

		if len(newItems) > 0 {
			newItemsString += "<b>Новые товары</b>\n\n"
			newItemsString += CatalogMessage(newItems)
		}

		if len(updateItems) > 0 {
			updateItemsString += "<b>Обновление цен</b>\n\n"
			updateItemsString += CatalogMessage(updateItems)
		}

		bot, err := tgbotapi.NewBotAPI(botToken)
		if err != nil {
			log.Panic(err)
		}

		users := database.GetUsers()

		for _, user := range users {
			if user.IsActive {
				if len(newItemsString) > 0 {
					msg := tgbotapi.NewMessage(user.Id, newItemsString)
					bot.Send(msg)
				}
				if len(updateItemsString) > 0 {
					msg := tgbotapi.NewMessage(user.Id, updateItemsString)
					bot.Send(msg)
				}
			}
		}
	}
}

func CatalogMessage(items []parsing.Item) string {
	var catalog string

	for _, val := range items {
		catalog += fmt.Sprintf("<a href=\"%s%s\">%s</a>\n", DNSDomain, val.Url, val.Title)
		catalog += fmt.Sprintf("<b>%d</b>", val.Price)
		if val.OldPrice != 0 {
			catalog += fmt.Sprintf(" <code>%d</code>", val.OldPrice)
		}
		if len(val.Desc) > 0 {
			catalog += fmt.Sprintf("<i>%s</i>", val.Desc)
		}
		catalog += "\n\n"
	}

	return catalog
}
