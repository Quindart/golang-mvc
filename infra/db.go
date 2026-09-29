package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type MyDB struct {
	Connector *sql.DB
}

func (myDB *MyDB) InitDB() {

	db, err := sql.Open("sqlite3", "./demo.db")

	if err != nil {
		fmt.Println(err)
		return
	}

	//
	fmt.Println("Connected to the SQLite database successfully.")
	myDB.Connector = db
	myDB.createUserTable()
}

func (myDB *MyDB) createUserTable() {

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		fullName TEXT,
		age INT,
		user_name TEXT,
		password TEXT
	)
	`
	_, err := myDB.Connector.Exec(query)

	if err != nil {
		fmt.Println("Error: create user table")
		return
	}
	fmt.Println("Success: create user table")

}
