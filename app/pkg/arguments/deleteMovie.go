package arguments

import (
	"database/sql"
	"flag"
	"fmt"
)

func DeleteMovie(db *sql.DB, arguments []string) {
	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	imdbIDToDelete := deleteCommand.String("imdbid", "tt0000001", "IMDb ID of a movie")
	deleteCommand.Parse(arguments[1:])

	query := "DELETE FROM movies WHERE IMDb_id = ?"
	result, err := db.Exec(query, *imdbIDToDelete)
	if err != nil {
		fmt.Println("Error deleting a movie:", err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("No movie found with the IMDb id:", *imdbIDToDelete)
	} else {
		fmt.Println("Movie deleted")
	}
}
