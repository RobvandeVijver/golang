package router

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"hz/package/models"
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
		if err := rows.Scan(&movie.IMDbID, &movie.Title, &movie.Rating, &movie.Year); err != nil {
			fmt.Println("Error scanning row:", err)
			c.Status(500)
			return
		}

		c.IndentedJSON(200, movie)
	}
}
