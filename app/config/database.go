package config

import (
	"database/sql"
	"fmt"
)

func StartDatabaseConnection(db *sql.DB) bool {
	var err error
	db, err = sql.Open("sqlite3", "./movies.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return true
	}
	return false
}
