package model

type UpdateItem struct {
	Item Item
	OldDiDiscountPrice int
}

type Item struct {
	ItemID string
	Title string
	Desc string
	URL string
	Price int
	OldPrice int
}

type User struct {
	ID int64
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

type Config struct {
	Debug *bool
	Logger *bool
	BotToken *string
	HoursUpdate *[]int
	DataSource *string
}