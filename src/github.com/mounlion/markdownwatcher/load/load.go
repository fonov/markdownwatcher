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

		"Accept": "*/*",
//"Accept-Encoding": "gzip, deflate, br",
"Accept-Language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-CN;q=0.6,zh;q=0.5",
"Connection": "keep-alive",
"Cookie": "_csrf=ea3482631f2c98783c0ebf5855bfdf9c25581faf1a65f2ce6188cbd8eee53b05a%3A2%3A%7Bi%3A0%3Bs%3A5%3A%22_csrf%22%3Bi%3A1%3Bs%3A32%3A%22826ehc3KWZVHSvV2SGZyg9zeUcMGrgQV%22%3B%7D; ipp_uid2=8gEcEZNBdgd8QmYW/PovKogT7S2odNZhd/BAn7g==; ipp_uid1=1527601972695; rerf=AAAAAFtDQvYRJ1uXAyfmAg==; PHPSESSID=a40ccfb6d1bfa7c3d74c8d98bdb9fd17; ipp_key=1531991507851/JFklO2d77mGld33Izq1JLg==; cartUserCookieIdent_v2=5fb454b6868f442ca46bbac0d5bade1a13f07e5934b7db4e4873c7089f340d00a%3A2%3A%7Bi%3A0%3Bs%3A22%3A%22cartUserCookieIdent_v2%22%3Bi%3A1%3Bs%3A36%3A%2214751f0b-2fc0-4940-8cfa-00eff9177910%22%3B%7D; phonesIdent=7d26948ce9c8deecdb4f467c922884b5acdf4aec97b2938b85facc7a7dca0ea4a%3A2%3A%7Bi%3A0%3Bs%3A11%3A%22phonesIdent%22%3Bi%3A1%3Bs%3A36%3A%221bc96025-6a04-4ccf-9a57-2f182cd33296%22%3B%7D; cto_lwid=8aea39cb-54f5-4d31-8d1d-a4625c4d43e1; _ga=GA1.2.1490900232.1532358827; _gid=GA1.2.727228129.1532358827; _ym_uid=1532358827123695750; _ym_d=1532358827; location-checker-user-identity-guid=e693e25e99cec2b47f982092187e1f9c07493a364937a4a1152270cee97fb9e3a%3A2%3A%7Bi%3A0%3Bs%3A35%3A%22location-checker-user-identity-guid%22%3Bi%3A1%3Bs%3A36%3A%22fd033d5a-a157-49da-a0c3-b3f92f021774%22%3B%7D; location-checker-getting-result=cc8376edd6b9db311c7b212b420f3093531bd8ac007860007cbcef24e64c6778a%3A2%3A%7Bi%3A0%3Bs%3A31%3A%22location-checker-getting-result%22%3Bi%3A1%3Bi%3A-1%3B%7D; tmr_detect=1%7C1532358827013; _gat=1; _ym_isad=1; _ym_visorc_7967056=b; cto_idcpy=c88300b1-61b3-46ec-9699-81aa881428e3",
"Host": "www.dns-shop.ru",
"Referer": "https://www.dns-shop.ru/catalog/markdown/",
"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1",
"X-CSRF-Token": "VQINTrLF1SRCZjzMAoDeMU_SaK9qeA7l0cde14dJ8QttMDsr2qbmbxU8aoRR9ogDHJUy1g1BdICEpBOQ9S6gXQ==",
"X-Requested-With": "XMLHttpRequest",
		//"X-Requested-With": "XMLHttpRequest",
		//"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36",
		//"Cookie": "***REMOVED***",
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

