package arguments

import (
	"database/sql"
	"fmt"
	"hz/pkg/models"
)

func ListMovies(db *sql.DB) {
	movies, shouldReturn := getMovies(db)
	if shouldReturn {
		return
	}

	for _, movie := range movies {
		fmt.Printf("%s\n", movie.Title)
	}
}

func getMovies(db *sql.DB) ([]movie.Movie, bool) {
	query := "SELECT * FROM movies"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error database connection:", err)
		return nil, true
	}
	defer rows.Close()

	var movies []movie.Movie
	for rows.Next() {
		var movie movie.Movie
		if err := rows.Scan(&movie.IMDbID, &movie.Title, &movie.Rating, &movie.Year); err != nil {
			fmt.Println("Error scanning the row", err)
			return nil, true
		}
		movies = append(movies, movie)
	}

	if len(movies) == 0 {
		fmt.Println("No movies found.")
		return nil, true
	}
	return movies, false
}
