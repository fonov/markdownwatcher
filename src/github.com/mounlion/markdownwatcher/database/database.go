package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/mounlion/markdownwatcher/parsing"
	"database/sql"
	"log"
)


func main ()  {
	var Items []parsing.Item
	Items = append(Items, parsing.Item{"Title", "Url", "sha256", "2222", 12, 1200})

	PrepareItems(Items)
}

func PrepareItems(items []parsing.Item)  {
	// получить список объектов по идексам в воде мы должны получит массив объектов ид цена
	// расчитывем какие новые а какие обновленные
	// insert новые объекты
	// update объекты для обновление
	// возврашаем новые и обновленные оъеты

	db, err := sql.Open("sqlite3", "./catalog.db")
	checkErr(err)
	defer db.Close()

	sqlStr := "select id, price from items when id in"
	var searchIDs string

	for _, val := range items {
		searchIDs += val.ItemId
	}

	sqlStr += "["+searchIDs+"]"

	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
	}

	var (
		id string
		price int
		updateItems []*parsing.Item
		newItems []*parsing.Item
	)

	for rows.Next() {
		err = rows.Scan(&id, &price)
		checkErr(err)
		for _, val := range items {
			if val.ItemId == id && val.Price != price {
				updateItems = append(updateItems, &val)
				break
			}
			if val.ItemId == id && val.Price == price { break }

			newItems = append(newItems, &val)
		}
	}

	rows.Close()

	insertStr := "INSERT INTO items(id, title, url, description, price, oldPrice) values "
	var insertVals string
	var insertData 

	for _, val := range newItems {
		insertVals += "(?,?,?,?,?,?),"
	}

	//rows, err := db.Query("CREATE TABLE family1 (member_id INT NOT NULL, name VARCHAR(50), relation VARCHAR(50));")

	//checkErr(err)

	//fmt.Println(rows)

	//db.Close()

	//sqlStr := "INSERT INTO items(id, title, url, description, price, oldPrice) values (?,?,?,?,?,?)"
	//vals := []interface{}{}
	//
	//items[0].title
	//
	//var item parsing.Item
	//
	//for _, row := range items {
	//	item = row
	//
	//	sqlStr += " (?,?,?,?,?,?),"
	//	//row.
	//	vals = append(vals, item.tile)
	//}
	////trim the last ,
	//sqlStr = sqlStr[0:len(sqlStr)-2]
	////prepare the statement
	//stmt, _ := db.Prepare(sqlStr)
	//
	////format all vals at once
	//res, _ := stmt.Exec(vals...)
	//
	//
	//db, err := sql.Open("sqlite3", "./catalog.db")
	//checkErr(err)
	//
	//defer db.Close()
	//
	stmt, err := db.Prepare("INSERT INTO items(id, title, url, description, price, oldPrice) values(?,?,?,?,?,?)")
	//checkErr(err)
	//
	//res, err := stmt.Exec()
	//checkErr(err)
	//
	//id, err := res.LastInsertId()
	//checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}