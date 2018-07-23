package main

import (
	"time"
	"github.com/mounlion/markdownwatcher/parsing"
	"github.com/mounlion/markdownwatcher/load"
	"github.com/mounlion/markdownwatcher/database"
	"github.com/mounlion/markdownwatcher/bot"
)

var (
	hoursUpdate = [...]int {9, 11, 12, 13, 14, 15, 16, 17, 18, 22, 23, 0}
)


func main() {
	go bot.Front()

	for {
		now := time.Now()
		for _, v := range hoursUpdate {
			if v == now.Hour() {
				html := load.Catalog()
				catalog := parsing.Catalog(html)
				newItems, updateItems := database.PrepareItems(catalog)
				bot.SendCatalog(newItems, updateItems)
				break
			}
		}
		time.Sleep(time.Hour)
	}
}