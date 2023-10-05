package arguments

import (
	"database/sql"
	"fmt"
	"hz/pkg/models"
)

func ListMovies(db *sql.DB) {
	query := "SELECT * FROM movies"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error database connection:", err)
		return
	}
	defer rows.Close()

	var movies []movie.Movie
	for rows.Next() {
		var movie movie.Movie
		if err := rows.Scan(&movie.IMDbID, &movie.Title, &movie.Rating, &movie.Year); err != nil {
			fmt.Println("Error scanning the row", err)
			return
		}
		movies = append(movies, movie)
	}

	if len(movies) == 0 {
		fmt.Println("No movies found.")
		return
	}

	for _, movie := range movies {
		fmt.Printf("%s\n", movie.Title)
	}
}
