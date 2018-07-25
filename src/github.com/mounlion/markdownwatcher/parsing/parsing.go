package parsing

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
	"strconv"
	"github.com/mounlion/markdownwatcher/model"
	"log"
)
var (
	Logger *bool
)

func SetInitialValue(_Logger *bool)  {
	Logger = _Logger
}

func Catalog(rawHtml string) []model.Item {

	var (
		typeOfWaitData int
		Items []model.Item
		isWaitText = false
		z = html.NewTokenizer(strings.NewReader(rawHtml))
		tempItem = model.Item{}
		isWaitTextList bool
		typeOfWaitListDate int 
		lastClass = make([]string, 3)
		countLBL int
	)

	if *Logger {log.Printf("Start parsing catalog")}

	for {
		tt := z.Next()
		if tt == html.ErrorToken { break }

		switch {
		case tt == html.StartTagToken:
			t := z.Token()
			for _, val := range t.Attr {
				if  val.Key == "class" {
					if len(lastClass[len(lastClass)-1]) == 0 {
						for i, v := range lastClass {
							if len(v) == 0 {
								lastClass[i] = val.Val
								break
							}
						}
					} else {
						lastClass = append(lastClass[1:], val.Val)
					}
				}
			}
			switch t.Data {
			case "div":
				for _, val := range t.Attr {
					switch {
					case val.Key == "data-id" && val.Val == "product":
						if len(tempItem.Title) >  0 {
							Items = append(Items, tempItem)
							tempItem = model.Item{}
						}
					case val.Key == "class" && val.Val == "item-name":
						isWaitText = true
						typeOfWaitData = 1
						break
					case val.Key == "class" && val.Val == "markdown-price-old":
						isWaitText = true
						typeOfWaitData = 3
						break
					case val.Key == "class" && val.Val == "characteristic-description":
						isWaitTextList = false
						break
					}
				}
			case "a":
				if len(tempItem.Url) > 0 { break }
				const className = "ec-price-item-link"
				var href string
				isNeedClass := false
				for _, val := range t.Attr {
					if val.Key == "href" {
						href = val.Val
					}
					if val.Key == "class" && val.Val == className {
						isNeedClass = true
					}
					if isNeedClass && len(href) > 0 { break }
				}
				if isNeedClass {
					tempItem.Url = href
					url := strings.Split(href, "/")
					tempItem.ItemId = url[3]
				}
			case "span":
				for _, val := range t.Attr {
					switch {
					case val.Key == "data-of" && val.Val == "price-total":
						isWaitText = true
						typeOfWaitData = 2
						break
					}
				}
			case "ul":
				for _, val := range t.Attr {
					if val.Key == "class" && val.Val == "list-unstyled markdown-reasons" {
						isWaitTextList = true
						typeOfWaitListDate = 1
						break
					}
				}
			}
		case tt == html.TextToken:
			switch {
			case lastClass[0] == "item-desc" && lastClass[1] == "small-screens" && lastClass[2] == "ec-price-item-link":
				t := z.Token()
				tempItem.Desc = t.Data
				break
			case isWaitText:
				t := z.Token()
				switch typeOfWaitData {
				case 1:
					tempItem.Title = t.Data
					break
				case 2:
					price, err := strconv.Atoi(strings.Replace(t.Data, ` `, "", -1))
					if err != nil {
						fmt.Println(err)
					}
					tempItem.Price = price
					break
				case 3:
					str := strings.Replace(t.Data, "\u00a0", "", -1)
					oldPrice, err := strconv.Atoi(strings.Replace(str, ` `, "", -1))
					if err != nil {
						fmt.Println(err)
					}
					tempItem.OldPrice = oldPrice
					break
				}
				isWaitText = false
			case isWaitTextList:
				t := z.Token()
				switch typeOfWaitListDate {
				case 1:
					if lastClass[2] == "lbl" && countLBL == 0{
						tempItem.Desc += fmt.Sprintf("%s: ", t.Data)
						countLBL++
					}
					if lastClass[2] == "reasons-inline" {
						tempItem.Desc += fmt.Sprintf("%s. ", t.Data)
						countLBL--
					}
					break
				}
			}
		}
	}

	if *Logger {log.Printf("Finish parsing catalog. Count: %d.", len(Items))}

	return Items
}
