package arguments

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
)

func AddMovie(db *sql.DB) {
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	imdbID := addCommand.String("imdbid", "tt0000001", "IMDb ID of a movie")
	title := addCommand.String("title", "Carmencita", "Title of a movie")
	year := addCommand.Int("year", 1894, "Year of release of the movie")
	rating := addCommand.Float64("rating", 5.7, "IMDb-rate of a movie")
	addCommand.Parse(os.Args[2:])

	query := "INSERT INTO movies (IMDb_id, Title, Rating, Year) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, *imdbID, *title, *rating, *year)
	if err != nil {
		fmt.Println("Error by adding the movie details:", err.Error())
		return
	}
	fmt.Printf("IMDb id: %s\nTitle: %s\nRating: %.1f\nYear: %d\n", *imdbID, *title, *rating, *year)
}
