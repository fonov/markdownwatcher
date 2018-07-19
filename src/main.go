package main

import (
	"time"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

func main() {

	//tick := time.Tick(time.Hour)

	//select {
	//case <-tick:
	//	fmt.Printf("hello, world\n%v %T", time.Now(), time.Now())
	//}

	loadCatalog()

	/*


	var hoursUpdate = []int{9, 12, 16, 17, 18}

	for {
		now := time.Now()

		for _, v := range hoursUpdate {
			if v == now.Hour() {
				loadCatalog()
				break
			}
		}

		time.Sleep(time.Hour)
	}
	*/
}

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

func loadCatalog() {
	//fmt.Printf("WORK")
	//var sprockets SprocketsResponse

	//response, err := netClient.Get("https://facebook.github.io/react-native/movies.json")

	//firstJson := fetchCatalog(0)
	//var html = firstJson.Html
	//lastProductIndex := firstJson.LastProductIndex
	//var countOfRequest = int(math.Round(float64(firstJson.FilteredProductsCount)/float64(lastProductIndex)))
	//const countItemInRequest  =
	//for i := 0; i < countOfRequest; i++ {
	//
	//}


	//fmt.Println(string(buf))

	//json.Unmarshal(buf, &res)
	//str := `{"title": "sdsd", "description": "sdd"}`

	//fmt.Println(jsonObj.LastProductIndex, resp.StatusCode, http.StatusText(resp.StatusCode))

	lastProductIndex := 0

	var html = ""

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

	parsingCatalog(html)

	//fmt.Println(headers["Cookie"])


	//if (json.)
	//fmt.Println(json)
	//for {
	//
	//}


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
	//fmt.Println(string(buf))
	return jsonObj
}

func parsingCatalog(html string)  {
	//fmt.Println(html)
}

func analyzeData()  {
	//var createItem, updateItem = make([]Item), make([]Item)

}

func getDataFromDB()  {
	
}

func setDateInDB()  {

}

func updateDataInDB()  {

}

func sendTelegramMessage()  {

}