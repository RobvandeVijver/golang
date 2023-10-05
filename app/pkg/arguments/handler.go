package arguments

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

func ArgumentHandler(db *sql.DB, baseURL string) {
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

		default:
			fmt.Println("invalid input")
		}
	} else {
		err := http.ListenAndServe(baseURL, nil)
		if err != nil {
			panic(err)
		}
	}
}
