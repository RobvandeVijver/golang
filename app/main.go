package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"hz/api/router"
	"hz/config"
	arguments2 "hz/pkg/arguments"
)

var db *sql.DB

func main() {
	baseURL := config.GetHost()

	var err error
	db, err = sql.Open("sqlite3", "./movies.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
	}
	defer db.Close()

	router.ApiHandler()

	arguments2.ArgumentHandler(db, baseURL)
}
