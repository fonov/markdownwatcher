package bot

import (
	"fmt"
	"github.com/mounlion/markdownwatcher/database"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"github.com/mounlion/markdownwatcher/model"
)

const DNSDomain = "https://www.dns-shop.ru"
var (
	BotToken *string
	Logger *bool
)

func SetInitialValue(_BotToken *string, _Logger *bool)  {
	BotToken = _BotToken
	Logger = _Logger
}

func SendCatalog(newItems []model.Item, updateItems []model.UpdateItem)  {
	bot, err := tgbotapi.NewBotAPI(*BotToken)
	if err != nil {log.Panic(err)}

	users := database.GetUsers()

	if *Logger {log.Printf("Send Catalog. newItems: %d, updateItems: %d.", len(newItems), len(updateItems))}

	if len(newItems) > 0 || len(updateItems) > 0 {

		var (
			newItemsStringList []string
			updateItemsStringList []string
			newItemsLastIndex = 0
			updateItemsLastIndex = 0
		)

		if len(newItems) > 0 {
			newItemsStringList = make([]string, 1)
			newItemsStringList[newItemsLastIndex] += "<b>Новые товары</b>\n\n"
			for _, val := range newItems {
				if len(newItemsStringList[newItemsLastIndex] + CatalogMessage(val, 0)) > 4096 {
					newItemsStringList = append(newItemsStringList, CatalogMessage(val, 0))
					newItemsLastIndex++
				} else {
					newItemsStringList[newItemsLastIndex] += CatalogMessage(val, 0)
				}
			}
		}

		if len(updateItems) > 0 {
			updateItemsStringList = make([]string, 1)
			updateItemsStringList[updateItemsLastIndex] += "<b>Обновление цен</b>\n\n"
			for _, val := range updateItems {
				if len(updateItemsStringList[updateItemsLastIndex] + CatalogMessage(val.Item, val.OldDiDiscountPrice)) > 4096 {
					updateItemsStringList = append(updateItemsStringList, CatalogMessage(val.Item, val.OldDiDiscountPrice))
					updateItemsLastIndex++
				} else {
					updateItemsStringList[updateItemsLastIndex] += CatalogMessage(val.Item, val.OldDiDiscountPrice)
				}
			}
		}

		for _, user := range users {
			if user.IsActive {
				if len(newItemsStringList) > 0 {
					for _, mess := range newItemsStringList {
						sendMessage(bot, &user, &mess, false)
					}
				}
				if len(updateItemsStringList) > 0 {
					for _, mess := range updateItemsStringList {
						sendMessage(bot, &user, &mess, false)
					}
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

func sendMessage(bot *tgbotapi.BotAPI, user *model.User, message *string, DisableNotification bool)  {
	msg := tgbotapi.NewMessage(user.Id, *message)
	msg.ParseMode = "HTML"
	msg.DisableWebPagePreview = true
	msg.DisableNotification = DisableNotification
	_, err := bot.Send(msg)
	if err != nil {
		if err.Error() == "Forbidden: bot was blocked by the user" {
			database.Subscribe(int(user.Id), false)
		} else {
			log.Print(err.Error())
		}
	}
}

func SendServiceMessage(text string)  {
	users := database.GetUsers()

	bot, err := tgbotapi.NewBotAPI(*BotToken)
	if err != nil {
		log.Panic(err)
	}

	if *Logger {log.Printf("Send service message")}

	for _, user := range users {
		if user.IsActive && user.IsAdmin {
			sendMessage(bot, &user, &text, false)
		}
	}
}

func CatalogMessage(item model.Item, OldDiDiscountPrice int)string {
	var catalog string

	catalog += fmt.Sprintf("<a href=\"%s%s\">%s</a>\n", DNSDomain, item.Url, item.Title)
	catalog += fmt.Sprintf("<b>%d₽</b>", item.Price)
	if item.OldPrice != 0 {
		profit := 100-(float64(item.Price)/float64(item.OldPrice)*100)
		catalog += fmt.Sprintf("		%d₽ %.1f%%", item.OldPrice, profit)
	}
	if OldDiDiscountPrice != 0 {
		catalog += fmt.Sprintf("\n<i>Переуценка на %d₽</i>", OldDiDiscountPrice-item.Price)
	}
	if len(item.Desc) > 0 {
		catalog += fmt.Sprintf("\n\n<i>%s</i>", item.Desc)
	}
	catalog += "\n\n"

	return catalog
}