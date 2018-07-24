package load

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"github.com/mounlion/markdownwatcher/bot"
)

var (
	headers = map[string]string{
		"X-Requested-With": "XMLHttpRequest",
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36",
		"Cookie": "***REMOVED***",
	}
)

type JsonObject struct {
	FiltersOptions   string      `json:"filtersOptions"`
	IsNextLoadAvailable   bool      `json:"isNextLoadAvailable"`
	IsNextLoadFinal   bool      `json:"isNextLoadFinal"`
	LastProductIndex   int      `json:"lastProductIndex"`
	FilteredProductsCount   int      `json:"filteredProductsCount"`
	Result   bool      `json:"result"`
	Html   string      `json:"html"`
}

func Catalog() string {
	var lastProductIndex, html = 0, ""

	for {
		fmt.Println("Fetch offset %v", lastProductIndex)
		result, statusCode := fetchCatalog(lastProductIndex)
		if statusCode != 200 {
			fmt.Println("Fetch failed. Status code: ", statusCode)
			message := fmt.Sprintf("<b>Обнаружена проблема</b>\n\nСтатус ответа сервера: <code>%d</code>", statusCode)
			bot.SendServiceMessage(message)
			break
		}
		html += result.Html
		if result.IsNextLoadAvailable {
			lastProductIndex = result.LastProductIndex
		} else {
			fmt.Println("All fetch end")
			break
		}
	}

	return html
}

func fetchCatalog (offset int)  (JsonObject, int)  {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", "https://www.dns-shop.ru/catalogMarkdown/category/update/?offset="+strconv.Itoa(offset), nil)
	if err != nil {
		fmt.Println("NewRequest error")
	}
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	resp, err := netClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Http request error")
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	jsonObj := JsonObject{}
	if resp.StatusCode == 200 {
		json.Unmarshal(buf, &jsonObj)
		return jsonObj, resp.StatusCode
	} else {
		fmt.Println(string(buf))
		return jsonObj, resp.StatusCode
	}
}

