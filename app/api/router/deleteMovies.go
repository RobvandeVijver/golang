package router

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DeleteMovie() func(c *gin.Context) {
	return func(c *gin.Context) {

		movieID := c.Param("id")

		var db *sql.DB
		var err error
		db, err = sql.Open("sqlite3", "./movies.db")
		if err != nil {
			c.Status(500)
		}
		defer db.Close()

		query := "DELETE FROM movies WHERE IMDb_id = ?"
		result, errorDelete := db.Exec(query, movieID)
		if errorDelete != nil {
			c.Status(500)
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			fmt.Println("No movie found with the IMDb id:", movieID)
		} else {
			fmt.Println("Movie deleted")
		}

		c.Status(204)
	}
}
