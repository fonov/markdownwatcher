package main

import (
"time"
"github.com/mounlion/markdownwatcher/parsing"
"github.com/mounlion/markdownwatcher/load"
"github.com/mounlion/markdownwatcher/database"
"github.com/mounlion/markdownwatcher/bot"
"log"
"github.com/mounlion/markdownwatcher/config"
)

const appVersion = 1.5

func main() {
	config.GetConfig()

	if *config.Config.Logger {log.Printf("Start Mark Down Watcher v. %.1f. Debug: %t, Logger: %t", appVersion, *config.Config.Debug, *config.Config.Logger)}

	go bot.Front()

	for {
		now := time.Now()
		for _, v := range *config.Config.HoursUpdate {
			if v == now.Hour() {
				if *config.Config.Logger {log.Printf("Start synchronizations catalog")}
				html := load.Catalog(0)
				if len(html) == 0 { break }
				catalog := parsing.Catalog(html)
				newItems, updateItems := database.PrepareItems(catalog)
				bot.SendCatalog(newItems, updateItems)
				if *config.Config.Logger {log.Printf("Finish synchronizations catalog")}
				break
			}
		}
		time.Sleep(time.Hour)
	}
}