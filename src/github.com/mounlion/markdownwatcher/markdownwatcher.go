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

const appVersion = 1.4

func main() {
	config.GetConfig()

	if *config.Config.Logger {log.Printf("Start Mark Down Watcher v. %.1f. Debug: %t, Logger: %t", appVersion, *config.Config.Debug, *config.Config.Logger)}

	go bot.Front()

	for {
		now := time.Now()
		for _, v := range *config.Config.HoursUpdate {
			if v == now.Hour() {
				// запускаем цицкл по словарю
				// получаем название, получаем ключ города
				for cityKey, cityName := range *config.Config.Cities {
					// вывводи в консоль какой город начал обрабатываться
					if *config.Config.Logger {log.Printf("Start synchronizations catalog. [%s]", cityName)}
					// передаем в каталог ключ города
					html := load.Catalog(0, cityKey)
					if len(html) == 0 {
						if *config.Config.Logger {
							log.Printf("HTML retutn null")
							break
						}
					}
					catalog := parsing.Catalog(html)
					newItems, updateItems := database.PrepareItems(catalog)
					// передаем в имя города
					bot.SendCatalog(newItems, updateItems, cityName)
					// печаем имя города
					if *config.Config.Logger {log.Printf("Finish synchronizations catalog. [%s]", cityName)}
				}

				break
			}
		}
		time.Sleep(time.Hour)
	}
}