package main

import (
	"time"
	"github.com/mounlion/markdownwatcher/load"
	"github.com/mounlion/markdownwatcher/parsing"
	"fmt"
)

var (
	hoursUpdate = [...]int {9, 12, 13, 14, 15, 16, 17, 18}
)


func main() {
	for {
		now := time.Now()
		for _, v := range hoursUpdate {
			if v == now.Hour() {
				html := load.Catalog()
				catalog := parsing.Catalog(html)
				fmt.Println(catalog)
				break
			}
		}
		time.Sleep(time.Hour)
	}
}