package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"hz/api/router"
	"hz/config"
	arguments2 "hz/pkg/arguments"
)

var db *sql.DB

func main() {
	baseURL := config.GetHost()

	if config.StartDatabaseConnection() {
		return
	}
	defer db.Close(db)

	router.ApiHandler()

	arguments2.ArgumentHandler(db, baseURL)
}
