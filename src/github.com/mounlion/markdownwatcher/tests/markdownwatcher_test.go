package tests

import (
	"testing"
	"github.com/mounlion/markdownwatcher/load"
	"github.com/mounlion/markdownwatcher/parsing"
	"github.com/mounlion/markdownwatcher/model"
	"github.com/mounlion/markdownwatcher/database"
	"github.com/mounlion/markdownwatcher/bot"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"log"
)

func TestMarkDownWatcher(t *testing.T) {
	Logger := false
	DataSourceName := "__DB_PATH__"
	BotToken := "__BOT_TOKEN"

	parsing.SetInitialValue(&Logger)
	load.SetInitialValue(&Logger)
	database.SetInitialValue(&DataSourceName, &Logger)
	bot.SetInitialValue(&BotToken, &Logger)

	html := load.Catalog(0)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	var testFetch = make([]bool, 3)

	doc.Find(".product").Each(func(i int, s *goquery.Selection) {
		s.Find(".thumbnail").Each(func(i int, s *goquery.Selection) {
			s.Find(".markdown-caption").Each(func(i int, s *goquery.Selection) {
				testFetch[0]=true
			})
			s.Find(".item-price").Each(func(i int, s *goquery.Selection) {
				testFetch[1]=true
			})
			s.Find(".content-info-column").Each(func(i int, s *goquery.Selection) {
				testFetch[2]=true
			})
		})
	})

	for _, val := range testFetch {
		if !val {
			t.Error("Recived non valid html")
		}
	}

	rawHTML := `<h4>Электрочайники</h4><div class="products products-list"><div class="node-block"><div class="product" data-id="product"><i class="hidden" data-product-param="brand" data-value="Galaxy"></i><i class="hidden" data-product-param="category" data-value="Электрочайники"></i><div class="thumbnail"><div class="item-name-mobile"><a href="/catalog/markdown/19cbe709-2980-4538-9617-9343ba50b7d1/" class="ec-price-item-link" data-product-param="name">Электрочайник Galaxy GL 0101 белый</a></div><div class="image"><a class="show-popover ec-price-item-link" href="/catalog/markdown/19cbe709-2980-4538-9617-9343ba50b7d1/" role="button" target="_blank" tab-index="0" data-html data-toggle="popover" data-placement="left" data-trigger="hover" data-content-target="#img-popover-8e8a1309-55f2-11e5-91a6-00155d03361b" data-container="#pop-abs-8e8a1309-55f2-11e5-91a6-00155d03361b"><picture><source type="image/webp" media="(min-width: 992px)" data-srcset="https://c.dns-shop.ru/thumb/st1/fit/60/60/32caa8bc5cd3e6215266443609f909a5/e91efc7c6f4639d55c62ad7ef2379ed899edee84a6d240fa3a587d53efdfb99d.jpg.webp"><source type="image/webp" media="(min-width: 768px)" data-srcset="https://c.dns-shop.ru/thumb/st1/fit/120/120/5a5cfcf7878fe8e50ac3b9c377fa80ae/e91efc7c6f4639d55c62ad7ef2379ed899edee84a6d240fa3a587d53efdfb99d.jpg.webp"><source type="image/webp" media="(max-width: 767px)" data-srcset="https://c.dns-shop.ru/thumb/st1/fit/90/90/d38830ee67e2c53b57685b3a54d0805f/e91efc7c6f4639d55c62ad7ef2379ed899edee84a6d240fa3a587d53efdfb99d.jpg.webp"><source media="(min-width: 992px)" data-srcset="https://c.dns-shop.ru/thumb/st1/fit/60/60/32caa8bc5cd3e6215266443609f909a5/e91efc7c6f4639d55c62ad7ef2379ed899edee84a6d240fa3a587d53efdfb99d.jpg"><source media="(min-width: 768px)" data-srcset="https://c.dns-shop.ru/thumb/st1/fit/120/120/5a5cfcf7878fe8e50ac3b9c377fa80ae/e91efc7c6f4639d55c62ad7ef2379ed899edee84a6d240fa3a587d53efdfb99d.jpg"><source media="(max-width: 767px)" data-srcset="https://c.dns-shop.ru/thumb/st1/fit/90/90/d38830ee67e2c53b57685b3a54d0805f/e91efc7c6f4639d55c62ad7ef2379ed899edee84a6d240fa3a587d53efdfb99d.jpg"><img alt="Электрочайник Galaxy GL 0101 белый" data-src="https://c.dns-shop.ru/thumb/st1/fit/60/60/32caa8bc5cd3e6215266443609f909a5/e91efc7c6f4639d55c62ad7ef2379ed899edee84a6d240fa3a587d53efdfb99d.jpg"></picture></a><div id="img-popover-8e8a1309-55f2-11e5-91a6-00155d03361b" class="popover" role="tooltip"><div class="popover-content img-popover"><picture><source type="image/webp" media="(min-width: 320px)" data-srcset="https://c.dns-shop.ru/thumb/st1/fit/190/190/3dffc8564adcb4937e8bcf2ff6ad4074/e91efc7c6f4639d55c62ad7ef2379ed899edee84a6d240fa3a587d53efdfb99d.jpg.webp"><source media="(min-width: 320px)" data-srcset="https://c.dns-shop.ru/thumb/st1/fit/190/190/3dffc8564adcb4937e8bcf2ff6ad4074/e91efc7c6f4639d55c62ad7ef2379ed899edee84a6d240fa3a587d53efdfb99d.jpg"><img alt="Электрочайник Galaxy GL 0101 белый" data-src="https://c.dns-shop.ru/thumb/st1/fit/190/190/3dffc8564adcb4937e8bcf2ff6ad4074/e91efc7c6f4639d55c62ad7ef2379ed899edee84a6d240fa3a587d53efdfb99d.jpg"></picture></div></div><div id="pop-abs-8e8a1309-55f2-11e5-91a6-00155d03361b" class="pop-abs"></div></div><div class="content-info-column"><div class="vertical-container"><div class="item-code" data-product-param="code">1030375</div><div class="markdown-status"><span class="markdown-dot active"></span><span class="markdown-dot active"></span><span class="markdown-dot active"></span><span class="markdown-dot active"></span><span class="markdown-dot "></span></div></div></div><div class="markdown-caption"><div class="item-name"><a href="/catalog/markdown/19cbe709-2980-4538-9617-9343ba50b7d1/" class="ec-price-item-link" data-product-param="name">Электрочайник Galaxy GL 0101 белый</a></div><div class="item-desc"><div class="small-screens"><div class="markdown-status"><span class="lbl">Оценка состояния:</span>&nbsp;<span class="markdown-dot active"></span><span class="markdown-dot active"></span><span class="markdown-dot active"></span><span class="markdown-dot active"></span><span class="markdown-dot "></span></div></div><ul class="list-unstyled markdown-reasons"><li><a href="/catalog/markdown/19cbe709-2980-4538-9617-9343ba50b7d1/" class="ec-price-item-link"><span class="lbl">Комплект</span>: <span class="reasons-inline">полный комплект</span></a></li><li><a href="/catalog/markdown/19cbe709-2980-4538-9617-9343ba50b7d1/" class="ec-price-item-link"><span class="lbl">Внешний вид</span>: <span class="reasons-inline">другое</span></a></li><li><a href="/catalog/markdown/19cbe709-2980-4538-9617-9343ba50b7d1/" class="ec-price-item-link"><span class="lbl">Вид</span>: <span class="reasons-inline">локальный ремонт</span></a></li></ul><div class="characteristic-description">[1.7 л, 2200 Вт, открытый нагревательный элемент, фильтр, материал корпуса - пластик]</div></div></div><div class="item-price"><div class="vertical-container product-item-price"><div class="price" data-id="cart-summary"><div class="price_g"><span data-of="price-total" data-product-param="price" data-value="250.00">250</span>&nbsp;<i class="rub-icon"></i></div></div></div></div><div class="characteristic-description-mobile">[1.7 л, 2200 Вт, открытый нагревательный элемент, фильтр, материал корпуса - пластик]</div><div class="buttons"><div itemprop="availability" href="http://schema.org/InStock" class="order-avail-wrap"><div class="avail-text"><span class="available">В наличии: </span><a class="show-popover pseudo-link" href="#" data-container="body" tabindex="0" role="button" data-toggle="popover" data-trigger="focus" data-content-target="#avail-19cbe709-2980-4538-9617-9343ba50b7d1" data-html="true" data-placement="bottom" data-template="&lt;div class=&quot;popover markdown-popup&quot; role=&quot;tooltip&quot;&gt;&lt;div class=&quot;arrow&quot;&gt;&lt;/div&gt;&lt;h3 class=&quot;popover-title&quot;&gt;&lt;/h3&gt;&lt;div class=&quot;popover-content&quot;&gt;&lt;/div&gt;&lt;/div&gt;" onclick="event.preventDefault(); $(this).focus();">DNS, Гипермолл...</a></div><div id="avail-19cbe709-2980-4538-9617-9343ba50b7d1" class="hide"><div class="order-avail-popup"><ul class="list-unstyled avail-items"><li><p><b>Магазин:</b><br>DNS, Гипермолл Горбушкин Двор</p><p><b>Адрес:</b><br>Багратионовский проезд, д. 7, корп. 3-<a href="/shops/moscow/7c2966-dns-gipermoll-gorbushkin-dvor/" target="_blank">показать на карте</a></p><p><b>Режим работы:</b><br>Пн-Вс с 10:00 до 21:00</p></li></ul></div></div></div></div></div></div></div></div>`

	catalogs := parsing.Catalog(rawHTML)

	if len(catalogs) > 0 {
		item := model.Item{
			ItemID: "19cbe709-2980-4538-9617-9343ba50b7d1",
			Title: "Электрочайник Galaxy GL 0101 белый",
			Desc: "Комплект: полный комплект. Внешний вид: другое. Вид: локальный ремонт. ",
			Url: "/catalog/markdown/19cbe709-2980-4538-9617-9343ba50b7d1/",
			Price: 250,
			OldPrice: 0,
		}

		if catalogs[0].ItemID != item.ItemID {t.Error("ItemId musb be equil")}
		if catalogs[0].Title != item.Title {t.Error("Title musb be equil")}
		if catalogs[0].Desc != item.Desc {t.Error("Desc musb be equil")}
		if catalogs[0].Price != item.Price {t.Error("Price musb be equil")}
		if catalogs[0].OldPrice != item.OldPrice {t.Error("OldPrice musb be equil")}
	} else {
		t.Error("Nul catalogs but must be 1")
	}

	newItems, updateItems := database.PrepareItems(catalogs)

	if len(newItems) != 0 {
		t.Error("newItems must be null")
	}
	if len(updateItems) != 0 {
		t.Error("updateItems must be null")
	}

	bot.SendCatalog(catalogs, updateItems)
}