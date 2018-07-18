package main

import (
	"time"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"encoding/xml"
	"bytes"
)

type html struct {
	Body body `xml:"body"`
}
type body struct {
	Content string `xml:"p"`
}

func main() {

	//tick := time.Tick(time.Hour)

	//select {
	//case <-tick:
	//	fmt.Printf("hello, world\n%v %T", time.Now(), time.Now())
	//}

	b := []byte(`<!DOCTYPE html>
<html>
    <head>
        <title>
            Title of the document
        </title>
    </head>
    <body>
        body content 
        <p>more content</p>
    </body>
</html>`)

	h := html{}
	err := xml.NewDecoder(bytes.NewBuffer(b)).Decode(&h)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(h.Body.Content)

	/*
	loadCatalog()

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
		"Cookie": "ipp_uid2=GrlMTVhT0TRMbDA2/BYCH5iAYReJHAWDY73yjEA==; ipp_uid1=1513617735271; _csrf=7126513c167a59bdb9938b0f3ecf11521de2cfb83d735853256dae816300c8d8a%3A2%3A%7Bi%3A0%3Bs%3A5%3A%22_csrf%22%3Bi%3A1%3Bs%3A32%3A%22AHN6YsoLAKXOEFWL6fnC-mnO1wQNLBk-%22%3B%7D; inptime0_5758_ru=0; rerf=AAAAAFsz6J6p0sjXAxqbAg==; city_path=moscow; current_path=229ab0dff2f4c4d7f0747a8270ba8c53efc7dab3709c232e0eaa175572cc12fba%3A2%3A%7Bi%3A0%3Bs%3A12%3A%22current_path%22%3Bi%3A1%3Bs%3A36%3A%2230b7c1f3-03fb-11dc-95ee-00151716f9f5%22%3B%7D; cartUserCookieIdent_v2=d32b45a266fef1b7336aff9a340ff1ce6083ee7618abcd4ad46aff4b697ce9a4a%3A2%3A%7Bi%3A0%3Bs%3A22%3A%22cartUserCookieIdent_v2%22%3Bi%3A1%3Bs%3A36%3A%223df8c9a9-64d9-4a9a-a5b5-f5813ff1724e%22%3B%7D; phonesIdent=8788830b636c6078f0f585015dd12a7d9739cd0262c6ff65cda6426c468cca31a%3A2%3A%7Bi%3A0%3Bs%3A11%3A%22phonesIdent%22%3Bi%3A1%3Bs%3A36%3A%22d4e040f4-e7e4-4298-b0ed-09ed53883b2b%22%3B%7D; _ga=GA1.2.116532446.1531254662; cto_lwid=1abc4889-8216-493a-97a7-930b04e5bd8e; _ym_uid=153125466219078692; _ym_d=1531254662; location-checker-user-identity-guid=a07da41e24f5ade833b6f56bfb58bf1c71da92b7b33f111cce1640b64a60c4a8a%3A2%3A%7Bi%3A0%3Bs%3A35%3A%22location-checker-user-identity-guid%22%3Bi%3A1%3Bs%3A36%3A%22d0e8674c-be48-4cb7-8c2b-f9cb1db9855c%22%3B%7D; location-checker-getting-result=cc8376edd6b9db311c7b212b420f3093531bd8ac007860007cbcef24e64c6778a%3A2%3A%7Bi%3A0%3Bs%3A31%3A%22location-checker-getting-result%22%3Bi%3A1%3Bi%3A-1%3B%7D; PHPSESSID=753cb6c8ef238ad1310064926f5c7a86; opinionComplainedIdentity=bdcc6957dcfdb4adbc318b6176f00f45cfe0270d8ba1dc71136b646904713fe2a%3A2%3A%7Bi%3A0%3Bs%3A25%3A%22opinionComplainedIdentity%22%3Bi%3A1%3Bs%3A20%3A%22KESiZi_IY6V3X8pa1YEY%22%3B%7D; _vi=9e5e3fcfa69d95fc4d4c89c72f9d463171cee623bccba47a2245d00a724850baa%3A2%3A%7Bi%3A0%3Bs%3A3%3A%22_vi%22%3Bi%3A1%3Bs%3A32%3A%226892d42d0141885ffda095a69b21a7dc%22%3B%7D; ipp_key=1531942687629/XFESyQDGjhS3WJLc8D8UGw==; _gid=GA1.2.1817850253.1531942689; _gat=1; _ym_visorc_7967056=w; tmr_detect=1%7C1531942689738; _ym_isad=1",
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
		//fmt.Println("Fetch offset %v", lastProductIndex)
		result := fetchCatalog(lastProductIndex)
		if result.IsNextLoadAvailable {
			lastProductIndex = result.LastProductIndex
			html += result.Html
		} else {
			//fmt.Println("All fetch end %v", len(html))
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

type Item struct {
	title, url, itemId, desc string
	price, oldPrice int
}

func parsingCatalog(html string)  {
	
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