package load

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
	"io/ioutil"
	"encoding/json"
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
		result := fetchCatalog(lastProductIndex)
		if result.IsNextLoadAvailable {
			lastProductIndex = result.LastProductIndex
			html += result.Html
		} else {
			fmt.Println("All fetch end %v", len(html))
			break
		}
	}

	return html
}

func fetchCatalog (offset int)  JsonObject {
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
	json.Unmarshal(buf, &jsonObj)
	return jsonObj
}

