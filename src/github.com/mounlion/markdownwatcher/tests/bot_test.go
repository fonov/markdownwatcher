package tests

import (
	"testing"
	"github.com/mounlion/markdownwatcher/bot"
	"os"
	"github.com/mounlion/markdownwatcher/database"
	"github.com/mounlion/markdownwatcher/model"
)

func TestSendMessage(t *testing.T) {

	newItems := []model.Item{
		//parsing.Item{"be40e256-898c-4480-bea5-fc86fd23b45d", "GPS навигатор DEXP Auriga DS503", "", "/catalog/markdown/be40e256-898c-4480-bea5-fc86fd23b45d/", 1200, 2000},
		//parsing.Item{"be40e256-898c-4480-bea5-fc86fd23b45d", "GPS навигатор DEXP Auriga DS503", "", "/catalog/markdown/be40e256-898c-4480-bea5-fc86fd23b45d/", 1200, 2000},
		//parsing.Item{"be40e256-898c-4480-bea5-fc86fd23b45d", "GPS навигатор DEXP Auriga DS503", "", "/catalog/markdown/be40e256-898c-4480-bea5-fc86fd23b45d/", 1200, 2000},
		//parsing.Item{"be40e256-898c-4480-bea5-fc86fd23b45d", "GPS навигатор DEXP Auriga DS503", "", "/catalog/markdown/be40e256-898c-4480-bea5-fc86fd23b45d/", 1200, 2000},
	}

	updateItems := []model.UpdateItem{
		model.UpdateItem{
			model.Item{
				"6f4e2f77-07ad-43ee-b7dd-86458589dc29",
				"GPS навигатор DEXP Auriga DS501",
				"Комплект: нет упаковки. Внешний вид: потертости. Вид: локальный ремонт.",
				"/catalog/markdown/6f4e2f77-07ad-43ee-b7dd-86458589dc29/", 1099,	2299,
			},
			1199,
		},

		//parsing.Item{
		//	"6f4e2f77-07ad-43ee-b7dd-86458589dc29",
		//	"GPS навигатор DEXP Auriga DS501",
		//	"","/catalog/markdown/6f4e2f77-07ad-43ee-b7dd-86458589dc29/", 1099,	2299,
		//
		//},
		//parsing.Item{
		//	"6f4e2f77-07ad-43ee-b7dd-86458589dc29",
		//	"GPS навигатор DEXP Auriga DS501",
		//	"","/catalog/markdown/6f4e2f77-07ad-43ee-b7dd-86458589dc29/", 1099,	2299,
		//
		//},
	}

	var BotToken = "***REMOVED***"
	var DataSourceName = os.Getenv("GOPATH")+"***REMOVED***"
	var Logger = false

	database.SetInitialValue(&DataSourceName, &Logger)
	bot.SetInitialValue(&BotToken, &Logger)

	bot.SendCatalog(newItems, updateItems)
}