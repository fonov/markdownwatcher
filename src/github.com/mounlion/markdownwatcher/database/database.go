package database

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"log"
	"github.com/mounlion/markdownwatcher/model"
	"github.com/mounlion/markdownwatcher/config"
)

func PrepareItems(items []model.Item) ([]model.Item, []model.UpdateItem) {
	db, err := sql.Open("sqlite3", *config.Config.DataSource)
	CheckErr(err)
	defer db.Close()

	selectStr := "select id, price from items where id in ("
	selectIds := make([]interface{}, len(items))

	for i, val := range items {
		selectStr += "?"
		selectIds[i] = val.ItemID
		if len(items)-1 != i {
			selectStr += ","
		} else {
			selectStr += ")"
		}
	}

	rows, err := db.Query(selectStr, selectIds...)
	if err != nil {log.Fatal(err)}

	if *config.Config.Logger {log.Printf("select id, price from items where id = []model.Item")}

	defer rows.Close()

	var (
		id string
		price int
		updateItems []model.UpdateItem
		CountRows int
	)

	for rows.Next() {
		err = rows.Scan(&id, &price)
		CheckErr(err)
		for i, item := range items {
			if item.ItemID == id {
				if item.Price != price {
					updateItems = append(updateItems, model.UpdateItem{Item: item, OldDiDiscountPrice: price})
				}
				items = append(items[:i], items[i+1:]...)
				break
			}
		}
		CountRows++
	}

	if len(items) > 0 || len(updateItems) > 0 {
		tx, err := db.Begin()
		if err != nil {log.Fatal(err)}

		if len(items) > 0 {
			stmt, err := tx.Prepare("INSERT OR IGNORE INTO items(id, title, url, description, price, oldPrice) values (?,?,?,?,?,?)")
			if err != nil {log.Fatal(err)}
			defer stmt.Close()
			for _, item := range items {
				_, err = stmt.Exec(item.ItemID, item.Title, item.URL, NullString(item.Desc), item.Price, NullInt(item.OldPrice))
				if err != nil {log.Fatal(err)}
			}
			if *config.Config.Logger {log.Printf("insert %d items", len(items))}
		}
		if len(updateItems) > 0 {
			stmt, err := tx.Prepare("UPDATE items set price=? where id=?")
			if err != nil {log.Fatal(err)}
			defer stmt.Close()
			for _, item := range updateItems {
				_, err = stmt.Exec(item.Item.Price, item.Item.ItemID)
				if err != nil {log.Fatal(err)}
			}
			if *config.Config.Logger {log.Printf("update %d items", len(updateItems))}
		}

		tx.Commit()
	}

	if CountRows > 0 {
		return items, updateItems
	}

	return []model.Item{}, updateItems
}

func GetUsers() []model.User {
	db, err := sql.Open("sqlite3", *config.Config.DataSource)
	CheckErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT id, isActive, isAdmin from users")
	if err != nil {log.Fatal(err)}
	if *config.Config.Logger {log.Printf("SELECT id, isActive, isAdmin from users")}
	defer rows.Close()

	var (
		Users []model.User
		id int64
		isActive bool
		isAdmin bool
	)

	for rows.Next() {
		err = rows.Scan(&id, &isActive, &isAdmin)
		CheckErr(err)
		Users = append(Users, model.User{id, isActive, isAdmin})
	}

	return Users
}

func Subscribe(userID int, isActive bool) {
	db, err := sql.Open("sqlite3", *config.Config.DataSource)
	CheckErr(err)
	defer db.Close()

	stmt, err := db.Prepare("update users set isActive=? where id=?")
	CheckErr(err)

	res, err := stmt.Exec(isActive, userID)
	CheckErr(err)

	n, err := res.RowsAffected()
	CheckErr(err)

	if *config.Config.Logger {log.Printf("update users set isActive=%t where id=%d. RowsAffected: %d.", isActive, userID, n)}

	if n == 0 {
		stmt, err = db.Prepare("insert into users(id, isActive) values (?,?)")
		CheckErr(err)

		_, err = stmt.Exec(userID, isActive)
		CheckErr(err)

		if *config.Config.Logger {log.Printf("insert into users(id, isActive) values (%d,%t)", n, userID, isActive)}
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