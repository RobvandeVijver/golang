package router

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"hz/config"
)

func ApiHandler() {
	router, done := routerSettings()
	if done {
		return
	}

	router.Use(checkNotFound)

	// PATHS
	router.GET("/movies", getMovies())
	router.GET("/movies/:id", GetMovieDetails())
	router.POST("/movies", PostMovies())
	router.DELETE("movies/:id", deleteMovie())

	startRouter(router)
}

func startRouter(router *gin.Engine) {
	baseURL := config.GetHost()
	errorRouter := router.Run(baseURL)
	if errorRouter != nil {
		return
	}
}

func routerSettings() (*gin.Engine, bool) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	return router, false
}

func checkNotFound(c *gin.Context) {
	c.Next()
	if c.Writer.Status() == 404 {
		c.Status(404)
	}
}

func deleteMovie() func(c *gin.Context) {
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
