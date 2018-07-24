package tests

import (
"testing"
	"github.com/mounlion/markdownwatcher/parsing"
	"github.com/mounlion/markdownwatcher/database"
	"github.com/mounlion/markdownwatcher/bot"
	"fmt"
)

func TestPrepareCatalog(t *testing.T) {

	catalog := []parsing.Item{
		parsing.Item{"be40e256-898c-4480-bea5-fc86fd23b45d-ld", "GPS навигатор DEXP Auriga DS503", "", "/catalog/markdown/be40e256-898c-4480-bea5-fc86fd23b45d/", 1200, 2000},
		parsing.Item{"3de723fd-bf83-4af8-a792-4a59bfe909b4", "GPS навигатор DEXP Auriga DS503", "", "/catalog/markdown/be40e256-898c-4480-bea5-fc86fd23b45d/", 1300, 2000},
		parsing.Item{"7000b444-087b-4d8b-a762-8ba56088a98f", "GPS навигатор DEXP Auriga DS503", "", "/catalog/markdown/be40e256-898c-4480-bea5-fc86fd23b45d/", 2890, 2000},
	}

	newItems, updateItems := database.PrepareItems(catalog)

	fmt.Println(newItems, updateItems)

	bot.SendCatalog(newItems, updateItems)
}
