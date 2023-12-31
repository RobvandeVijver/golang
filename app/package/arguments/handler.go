package arguments

import (
	"database/sql"
	"fmt"
	"hz/api/router"
	"os"
)

func ArgumentHandler(db *sql.DB) bool {
	arguments := os.Args[1:]

	if len(arguments) > 0 {
		switch arguments[0] {
		case "help":
			PrintHelpMessage()

		case "add":
			AddMovie(db)

		case "list":
			ListMovies(db)

		case "details":
			GetMovieDetails(db, arguments)

		case "delete":
			DeleteMovie(db, arguments)

		case "summaries":
			router.ApiRequest(db)

		default:
			fmt.Println("invalid input")
			return false
		}
		return true
	}
	return false
}
