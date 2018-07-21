package bot

import (
	"github.com/mounlion/markdownwatcher/parsing"
	"fmt"
	"database/sql"
	"github.com/mounlion/markdownwatcher/database"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

const DNSDomain = "https://www.dns-shop.ru"

func main ()  {

}

func SendMessage(newItems []parsing.Item, updateItems[]parsing.Item)  {
	if len(newItems) > 0 || len(updateItems) > 0 {
		db, err := sql.Open("sqlite3", "./MarkDownWatcher.db")
		database.CheckErr(err)
		defer db.Close()

		rows, err := db.Query("SELECT id from users when isActive=?", true)

		var id int64

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

		bot, err := tgbotapi.NewBotAPI("***REMOVED***")
		if err != nil {
			log.Panic(err)
		}

		for rows.Next() {
			err = rows.Scan(&id)
			database.CheckErr(err)

			if len(newItemsString) > 0 {
				msg := tgbotapi.NewMessage(id, newItemsString)
				bot.Send(msg)
			}
			if len(updateItemsString) > 0 {
				msg := tgbotapi.NewMessage(id, updateItemsString)
				bot.Send(msg)
			}
		}

		rows.Close()
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
