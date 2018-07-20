package main

import (
	"time"
	"fmt"
	"github.com/mounlion/markdownwatcher/parsing"
)

var (
	hoursUpdate = [...]int {9, 12, 13, 14, 15, 16, 17, 18}
)


func main() {
	for {
		now := time.Now()
		for _, v := range hoursUpdate {
			if v == now.Hour() {
				//html := load.Catalog()
				//catalog := parsing.Catalog(html)

				//var Items []parsing.Item
				//Items = append(Items, parsing.Item{"Title", "Url", "sha256", "2222", 12, 1200})

				fmt.Println(parsing.Item{})
				break
			}
		}
		time.Sleep(time.Hour)
	}
}