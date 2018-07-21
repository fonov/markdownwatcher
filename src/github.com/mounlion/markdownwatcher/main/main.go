package main

import (
	"time"
	"fmt"
	"github.com/mounlion/markdownwatcher/parsing"
	"github.com/mounlion/markdownwatcher/load"
	"github.com/mounlion/markdownwatcher/database"
)

var (
	hoursUpdate = [...]int {9, 11, 12, 13, 14, 15, 16, 17, 18, 22, 23, 0}
)


func main() {
	for {
		now := time.Now()
		for _, v := range hoursUpdate {
			if v == now.Hour() {
				html := load.Catalog()
				catalog := parsing.Catalog(html)
				newItems, updateItems := database.PrepareItems(catalog)
				fmt.Println(newItems, updateItems)
				break
			}
		}
		time.Sleep(time.Hour)
	}
}