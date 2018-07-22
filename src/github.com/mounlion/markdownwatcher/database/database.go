package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/mounlion/markdownwatcher/parsing"
	"database/sql"
	"log"
)

const (
	DataSourceName = "./MarkDownWatcher.db"
)

func PrepareItems(items []parsing.Item) ([]parsing.Item, []parsing.Item) {
	db, err := sql.Open("sqlite3", DataSourceName)
	CheckErr(err)
	defer db.Close()

	selectStr := "select id, price from items where id in ("
	selectIds := make([]interface{}, len(items))

	for i, val := range items {
		selectStr += "?"
		selectIds[i] = val.ItemId
		if len(items)-1 != i {
			selectStr += ","
		} else {
			selectStr += ")"
		}
	}

	rows, err := db.Query(selectStr, selectIds...)
	if err != nil {log.Fatal(err)}
	defer rows.Close()

	var (
		id string
		price int
		updateItems []parsing.Item
		CountRows int
	)

	for rows.Next() {
		err = rows.Scan(&id, &price)
		CheckErr(err)
		for i, item := range items {
			if item.ItemId == id {
				if item.Price != price {
					updateItems = append(updateItems, item)
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
				_, err = stmt.Exec(item.ItemId, item.Title, item.Url, NullString(item.Desc), item.Price, NullInt(item.OldPrice))
				if err != nil {log.Fatal(err)}
			}
		}
		if len(updateItems) > 0 {
			stmt, err := tx.Prepare("UPDATE items set price=? where id=?")
			if err != nil {log.Fatal(err)}
			defer stmt.Close()
			for _, item := range updateItems {
				_, err = stmt.Exec(item.Price, item.ItemId)
				if err != nil {log.Fatal(err)}
			}
		}

		tx.Commit()
	}

	if CountRows > 0 {
		return items, updateItems
	} else {
		return []parsing.Item{}, []parsing.Item{}
	}
}

type User struct {
	Id int64
	IsActive bool
	IsAdmin bool
}

func GetUsers() []User {
	db, err := sql.Open("sqlite3", DataSourceName)
	CheckErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT id, isActive, isAdmin from users")
	if err != nil {log.Fatal(err)}
	defer rows.Close()

	var (
		Users []User
		id int64
		isActive bool
		isAdmin bool
	)

	for rows.Next() {
		err = rows.Scan(&id, &isActive, &isAdmin)
		CheckErr(err)
		Users = append(Users, User{id, isActive, isAdmin})
	}

	return Users
}

func Subscribe(userId int, isActive bool) {
	db, err := sql.Open("sqlite3", DataSourceName)
	CheckErr(err)
	defer db.Close()

	stmt, err := db.Prepare("update users set isActive=? where id=?")
	CheckErr(err)

	res, err := stmt.Exec(isActive, userId)
	CheckErr(err)

	n, err := res.RowsAffected()
	CheckErr(err)

	if n == 0 {
		stmt, err = db.Prepare("insert into users(id, isActive) values (?,?)")
		CheckErr(err)

		_, err = stmt.Exec(userId, isActive)
		CheckErr(err)
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