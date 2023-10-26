package router

import (
	"database/sql"
	"fmt"
	movie "hz/package/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMovies() func(c *gin.Context) {
	return func(c *gin.Context) {
		var db *sql.DB
		var err error
		db, err = sql.Open("sqlite3", "./movies.db")
		if err != nil {
			fmt.Println("Error opening database:", err)
		}
		defer db.Close()

		query := "SELECT IMDb_id, Title, Rating, Year FROM movies"
		rows, errorQuery := db.Query(query)
		if errorQuery != nil {
			fmt.Println("Error database connection:", errorQuery)
			return
		}
		defer rows.Close()

		var movies []movie.Movie
		for rows.Next() {
			var movieInfo movie.Movie
			if err := rows.Scan(&movieInfo.IMDbID, &movieInfo.Title, &movieInfo.Rating, &movieInfo.Year); err != nil {
				fmt.Println("Error scanning the row", err)
				return
			}
			movies = append(movies, movieInfo)
		}

		var simplifiedMovies []struct {
			IMDbID *string  `json:"imdb_id"`
			Title  *string  `json:"title"`
			Rating *float64 `json:"rating"`
			Year   *string  `json:"year"`
		}

		for _, m := range movies {
			simplifiedMovies = append(simplifiedMovies, struct {
				IMDbID *string  `json:"imdb_id"`
				Title  *string  `json:"title"`
				Rating *float64 `json:"rating"`
				Year   *string  `json:"year"`
			}{
				IMDbID: m.IMDbID,
				Title:  m.Title,
				Rating: m.Rating,
				Year:   m.Year,
			})
		}

		c.IndentedJSON(http.StatusOK, simplifiedMovies)
	}
}
