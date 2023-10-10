package arguments

import (
	"database/sql"
	"fmt"
	"hz/package/models"
)

func GetMovieDetails(db *sql.DB, arguments []string) {
	if len(arguments) < 2 {
		fmt.Println("Use: movie details <IMDb ID>")
		return
	}
	imdbID := arguments[2]
	query := "SELECT * FROM movies WHERE IMDb_id = ?"
	rows, err := db.Query(query, imdbID)
	if err != nil {
		fmt.Println("Error accessing movie details:", err.Error())
		return
	}
	defer rows.Close()

	if !rows.Next() {
		fmt.Printf("No movie found with the IMDb id: %s\n", imdbID)
		return
	}

	var movie movie.Movie
	if err := rows.Scan(&movie.IMDbID, &movie.Title, &movie.Rating, &movie.Year); err != nil {
		fmt.Println("Error scanning row:", err)
		return
	}
	fmt.Printf("IMDb id: %s\nTitle: %s\nRating: %.1f\nYear: %d\n", movie.IMDbID, movie.Title, movie.Rating, movie.Year)
}
