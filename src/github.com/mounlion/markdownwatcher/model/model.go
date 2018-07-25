package model

import "github.com/mounlion/markdownwatcher/parsing"

type UpdateItem struct {
	Item parsing.Item
	OldDiDiscountPrice int
}
