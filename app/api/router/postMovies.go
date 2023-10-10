package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"hz/package/models"
	"net/http"
)

func PostMovies() func(c *gin.Context) {
	return func(c *gin.Context) {

		var db *sql.DB
		var err error
		db, err = sql.Open("sqlite3", "./movies.db")
		if err != nil {
			c.Status(500)
		}
		defer db.Close()

		var newMovie movie.Movie

		if err := c.BindJSON(&newMovie); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		insertStatement := "INSERT INTO movies (IMDb_id, Title, Rating, Year) VALUES (?, ?, ?, ?)"
		_, error := db.Exec(insertStatement, newMovie.IMDbID, newMovie.Title, newMovie.Rating, newMovie.Year)
		if error != nil {
			c.Status(500)
			return
		}

		c.IndentedJSON(201, newMovie)
	}
}
