package tests

import (
	"testing"
	"github.com/mounlion/markdownwatcher/parsing"
	"github.com/mounlion/markdownwatcher/bot"
)

func TestSendMessage(t *testing.T) {

	newItems := []parsing.Item{
		parsing.Item{"be40e256-898c-4480-bea5-fc86fd23b45d", "GPS навигатор DEXP Auriga DS503", "", "/catalog/markdown/be40e256-898c-4480-bea5-fc86fd23b45d/", 1200, 2000},
		parsing.Item{"be40e256-898c-4480-bea5-fc86fd23b45d", "GPS навигатор DEXP Auriga DS503", "", "/catalog/markdown/be40e256-898c-4480-bea5-fc86fd23b45d/", 1200, 2000},
		parsing.Item{"be40e256-898c-4480-bea5-fc86fd23b45d", "GPS навигатор DEXP Auriga DS503", "", "/catalog/markdown/be40e256-898c-4480-bea5-fc86fd23b45d/", 1200, 2000},
		parsing.Item{"be40e256-898c-4480-bea5-fc86fd23b45d", "GPS навигатор DEXP Auriga DS503", "", "/catalog/markdown/be40e256-898c-4480-bea5-fc86fd23b45d/", 1200, 2000},
	}

	updateItems := []parsing.Item{
		parsing.Item{
			"6f4e2f77-07ad-43ee-b7dd-86458589dc29",
			"GPS навигатор DEXP Auriga DS501",
			"","/catalog/markdown/6f4e2f77-07ad-43ee-b7dd-86458589dc29/", 1099,	2299,

		},
		parsing.Item{
			"6f4e2f77-07ad-43ee-b7dd-86458589dc29",
			"GPS навигатор DEXP Auriga DS501",
			"","/catalog/markdown/6f4e2f77-07ad-43ee-b7dd-86458589dc29/", 1099,	2299,

		},
		parsing.Item{
			"6f4e2f77-07ad-43ee-b7dd-86458589dc29",
			"GPS навигатор DEXP Auriga DS501",
			"","/catalog/markdown/6f4e2f77-07ad-43ee-b7dd-86458589dc29/", 1099,	2299,

		},
	}

	bot.SendCatalog(newItems, updateItems)
}