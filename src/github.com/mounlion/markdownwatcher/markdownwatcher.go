package main

import (
	"time"
	"github.com/mounlion/markdownwatcher/parsing"
	"github.com/mounlion/markdownwatcher/load"
	"github.com/mounlion/markdownwatcher/database"
	"github.com/mounlion/markdownwatcher/bot"
	"flag"
	"log"
)

var (
	HoursUpdate []int
	DataSourceName string
	BotToken string
)


func main() {
	Debug := flag.Bool("debug", false, "Use debug mode for create and update Mark Down Watcher")
	Logger := flag.Bool("log", false, "Use log for view all processes")
	flag.Parse()

	if *Logger {log.Printf("Start Mark Down Watcher. Debug: %t, Logger: %t", *Debug, *Logger)}

	if *Debug {
		BotToken = "***REMOVED***"
		HoursUpdate = append(HoursUpdate,  9, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 0)
		DataSourceName = "***REMOVED***"
	} else {
		BotToken = "***REMOVED***"
		HoursUpdate = append(HoursUpdate,  8, 10, 12, 14, 17, 19, 22, 23)
		DataSourceName = "/home/fonov/markdownwatcher/MarkDownWatcher.prod.db"
	}

	database.SetInitialValue(&DataSourceName, Logger)
	bot.SetInitialValue(&BotToken, Logger)
	load.SetInitialValue(Logger)
	parsing.SetInitialValue(Logger)

	go bot.Front(BotToken)

	for {
		now := time.Now()
		log.Println(now)
		for _, v := range HoursUpdate {
			if v == now.Hour() {
				if *Logger {log.Printf("Start synchronizations catalog")}
				html := load.Catalog(0)
				if len(html) == 0 { break }
				catalog := parsing.Catalog(html)
				newItems, updateItems := database.PrepareItems(catalog)
				bot.SendCatalog(newItems, updateItems)
				if *Logger {log.Printf("Finish synchronizations catalog")}
				break
			}
		}
		time.Sleep(time.Hour)
	}
}