package model

type UpdateItem struct {
	Item Item
	OldDiDiscountPrice int
}

type Item struct {
	ItemId string
	Title string
	Desc string
	Url string
	Price int
	OldPrice int
}

type User struct {
	Id int64
	IsActive bool
	IsAdmin bool
}

type JsonObject struct {
	FiltersOptions   string      `json:"filtersOptions"`
	IsNextLoadAvailable   bool      `json:"isNextLoadAvailable"`
	IsNextLoadFinal   bool      `json:"isNextLoadFinal"`
	LastProductIndex   int      `json:"lastProductIndex"`
	FilteredProductsCount   int      `json:"filteredProductsCount"`
	Result   bool      `json:"result"`
	Html   string      `json:"html"`
}