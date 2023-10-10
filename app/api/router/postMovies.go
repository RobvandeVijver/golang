package router

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"hz/package/models"
	"net/http"
)

// REST
func PostMovies() func(c *gin.Context) {
	return func(c *gin.Context) {

		var db *sql.DB
		var err error
		db, err = sql.Open("sqlite3", "./movies.db")
		if err != nil {
			fmt.Println("Error opening database:", err)
		}
		defer db.Close()

		var newMovie movie.Movie

		if err := c.BindJSON(&newMovie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		insertStatement := "INSERT INTO movies (IMDb_id, Title, Rating, Year) VALUES (?, ?, ?, ?)"
		_, error := db.Exec(insertStatement, newMovie.IMDbID, newMovie.Title, newMovie.Rating, newMovie.Year)
		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
			return
		}

		c.IndentedJSON(201, newMovie)
	}
}
