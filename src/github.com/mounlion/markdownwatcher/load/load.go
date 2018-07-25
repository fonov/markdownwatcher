package load

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"github.com/mounlion/markdownwatcher/bot"
	"github.com/mounlion/markdownwatcher/model"
	"log"
)

var (
	headers = map[string]string{
		"X-Requested-With": "XMLHttpRequest",
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36",
		"Cookie": "***REMOVED***",
	}
	Logger *bool
)

func SetInitialValue(_Logger *bool)  {
	Logger = _Logger
}

func Catalog() string {
	var lastProductIndex, html = 0, ""

	for {
		if *Logger {log.Printf("Fetch offset %d", lastProductIndex)}
		result, statusCode := fetchCatalog(lastProductIndex)
		if statusCode != 200 {
			if *Logger {log.Printf("Fetch failed. Status code: %d", statusCode)}
			message := fmt.Sprintf("<b>Обнаружена проблема</b>\n\nСтатус ответа сервера: <code>%d</code>", statusCode)
			bot.SendServiceMessage(message)
			break
		}
		html += result.Html
		if result.IsNextLoadAvailable {
			lastProductIndex = result.LastProductIndex
		} else {
			if *Logger {log.Printf("All fetch end")}
			break
		}
	}

	return html
}

func fetchCatalog (offset int)  (model.JsonObject, int)  {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", "https://www.dns-shop.ru/catalogMarkdown/category/update/?offset="+strconv.Itoa(offset), nil)
	if err != nil {log.Printf("NewRequest error")}
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	resp, err := netClient.Do(req)
	defer resp.Body.Close()
	if err != nil {log.Printf("Http request error")}
	buf, _ := ioutil.ReadAll(resp.Body)
	jsonObj := model.JsonObject{}
	if resp.StatusCode == 200 {
		json.Unmarshal(buf, &jsonObj)
		return jsonObj, resp.StatusCode
	} else {
		fmt.Println(string(buf))
		return jsonObj, resp.StatusCode
	}
}

