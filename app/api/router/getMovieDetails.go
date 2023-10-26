package router

import (
	"database/sql"
	"fmt"
	movie "hz/package/models"

	"github.com/gin-gonic/gin"
)

func GetMovieDetails() func(c *gin.Context) {
	return func(c *gin.Context) {
		movieID := c.Param("id")

		var db *sql.DB
		var err error
		db, err = sql.Open("sqlite3", "./movies.db")
		if err != nil {
			c.Status(500)
		}
		defer db.Close()

		query := "SELECT * FROM movies WHERE IMDb_id = ?"
		rows, err := db.Query(query, movieID)
		if err != nil {
			c.Status(500)
			return
		}
		defer rows.Close()

		if !rows.Next() {
			fmt.Printf("No movie found with the IMDb id: %s\n", movieID)
			c.Status(404)
			return
		}

		var movie movie.Movie
		if err := rows.Scan(&movie.IMDbID, &movie.Title, &movie.Rating, &movie.Year, &movie.Plot); err != nil {
			fmt.Println("Error scanning row:", err)
			c.Status(500)
			return
		}

		var simplifiedMovie []struct {
			IMDbID *string  `json:"imdb_id"`
			Title  *string  `json:"title"`
			Rating *float64 `json:"rating"`
			Year   *string  `json:"year"`
		}

		simplifiedMovie = append(simplifiedMovie, struct {
			IMDbID *string  `json:"imdb_id"`
			Title  *string  `json:"title"`
			Rating *float64 `json:"rating"`
			Year   *string  `json:"year"`
		}{
			IMDbID: movie.IMDbID,
			Title:  movie.Title,
			Rating: movie.Rating,
			Year:   movie.Year,
		})

		c.IndentedJSON(200, simplifiedMovie)
	}
}
