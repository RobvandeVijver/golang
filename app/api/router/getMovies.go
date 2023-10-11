package router

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"hz/package/models"
	"net/http"
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

		query := "SELECT * FROM movies"
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

		c.IndentedJSON(http.StatusOK, movies)
	}
}
