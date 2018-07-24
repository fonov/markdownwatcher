package main

import (
	"time"
	"github.com/mounlion/markdownwatcher/parsing"
	"github.com/mounlion/markdownwatcher/load"
	"github.com/mounlion/markdownwatcher/database"
	"github.com/mounlion/markdownwatcher/bot"
	"flag"
	"os"
)

var (
	HoursUpdate []int
	DataSourceName string
	BotToken string
)


func main() {
	Debug := flag.Bool("debug", false, "Use debug mode for create and update Mark Down Watcher")
	flag.Parse()

	if *Debug {
		BotToken = "***REMOVED***"
		HoursUpdate = append(HoursUpdate,  9, 11, 12, 13, 14, 15, 16, 17, 18, 22, 23, 0)
		DataSourceName = os.Getenv("GOPATH")+"***REMOVED***"
	} else {
		BotToken = "***REMOVED***"
		HoursUpdate = append(HoursUpdate,  8, 10, 12, 14, 17, 19, 22)
		DataSourceName = "/home/fonov/markdownwatcher/MarkDownWatcher.prod.db"
	}

	database.SetDataSourceName(&DataSourceName)
	bot.SetBotToken(&BotToken)

	go bot.Front(BotToken)

	for {
		now := time.Now()
		for _, v := range HoursUpdate {
			if v == now.Hour() {
				html := load.Catalog()
				if len(html) == 0 { break }
				catalog := parsing.Catalog(html)
				newItems, updateItems := database.PrepareItems(catalog)
				bot.SendCatalog(newItems, updateItems)
				break
			}
		}
		time.Sleep(time.Hour)
	}
}