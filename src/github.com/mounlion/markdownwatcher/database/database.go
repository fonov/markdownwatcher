package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/mounlion/markdownwatcher/parsing"
	"database/sql"
	"log"
)


func main () {
	var Items []parsing.Item
	Items = append(Items, parsing.Item{"Title", "Url", "sha256", "2222", 12, 1200})

	PrepareItems(Items)
}

type CatalogData struct {
	Id string
	Price int
}

func PrepareItems(items []parsing.Item) ([]parsing.Item, []parsing.Item) {
	// получить список объектов по идексам в воде мы должны получит массив объектов ид цена
	// расчитывем какие новые а какие обновленные
	// insert новые объекты
	// update объекты для обновление
	// возврашаем новые и обновленные оъеты

	db, err := sql.Open("sqlite3", "./MarkDownWatcher.db")
	CheckErr(err)
	defer db.Close()

	sqlStr := "select id, price from items where id in ("

	for i, val := range items {
		sqlStr += `"`+val.ItemId+`"`
		if len(items)-1 != i {
			sqlStr += ","
		} else {
			sqlStr += ")"
		}
	}

	rows, err := db.Query(sqlStr)
	if err != nil {log.Fatal(err)}
	defer rows.Close()

	var (
		id string
		price int
		catalogsData []CatalogData
		updateItems []parsing.Item
		newItems []parsing.Item
		isFirstRun = false
	)

	for rows.Next() {
		err = rows.Scan(&id, &price)
		CheckErr(err)
		catalogsData = append(catalogsData, CatalogData{id, price})
	}

	if len(catalogsData) > 0 {
		for _, val := range items {
			isNewItem := true
			for _, dbVal := range catalogsData {
				if val.ItemId == dbVal.Id && val.Price == dbVal.Price {
					isNewItem = false
					break
				}
				if val.ItemId == dbVal.Id && val.Price != dbVal.Price {
					isNewItem = false
					updateItems = append(updateItems, val)
					break
				}
			}

			if isNewItem {
				newItems = append(newItems, val)
			}
		}
	} else {
		isFirstRun = true
		newItems = items
	}

	if len(newItems) > 0 || len(updateItems) > 0 {
		tx, err := db.Begin()
		if err != nil {log.Fatal(err)}

		if len(newItems) > 0 {
			stmt, err := tx.Prepare("INSERT OR IGNORE INTO items(id, title, url, description, price, oldPrice) values (?,?,?,?,?,?)")
			if err != nil {log.Fatal(err)}
			defer stmt.Close()
			for _, item := range newItems {
				_, err = stmt.Exec(item.ItemId, item.Title, item.Url, NullString(item.Desc), item.Price, NullInt(item.OldPrice))
				if err != nil {log.Fatal(err)}
			}
		}
		if len(updateItems) > 0 {
			stmt, err := tx.Prepare("UPDATE items set price=? where id=?")
			if err != nil {log.Fatal(err)}
			defer stmt.Close()
			for _, item := range newItems {
				_, err = stmt.Exec(item.Price, item.ItemId)
				if err != nil {log.Fatal(err)}
			}
		}

		tx.Commit()
	}

	if isFirstRun {
		return []parsing.Item{}, []parsing.Item{}
	} else {
		return newItems, updateItems
	}
}

func NullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid: true,
	}
}

func NullInt(i int) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: int64(i),
		Valid: true,
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}