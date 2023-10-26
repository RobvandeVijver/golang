package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"hz/api/router"
	arguments2 "hz/package/arguments"
)

var db *sql.DB

func main() {

	var err error
	db, err = sql.Open("sqlite3", "./movies.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
	}
	defer db.Close()

	isCommand := arguments2.ArgumentHandler(db)
	if !isCommand {
		router.ApiHandler()
	}
}
