package router

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"hz/pkg/models"
)

func GetMovieDetails() func(c *gin.Context) {
	return func(c *gin.Context) {
		movieID := c.Param("id")

		var db *sql.DB
		var err error
		db, err = sql.Open("sqlite3", "./movies.db")
		if err != nil {
			fmt.Println("Error opening database:", err)
		}
		defer db.Close()

		query := "SELECT * FROM movies WHERE IMDb_id = ?"
		rows, err := db.Query(query, movieID)
		if err != nil {
			fmt.Println("Error accessing movie details:", err.Error())
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}
		defer rows.Close()

		if !rows.Next() {
			fmt.Printf("No movie found with the IMDb id: %s\n", movieID)
			c.JSON(404, gin.H{"error": "Movie not found"})
			return
		}

		var movie movie.Movie
		if err := rows.Scan(&movie.IMDbID, &movie.Title, &movie.Rating, &movie.Year); err != nil {
			fmt.Println("Error scanning row:", err)
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		c.JSON(200, movie)
	}
}
