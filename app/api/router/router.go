package router

import (
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
	router.PUT("/movies/:id", putMovie())
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
	router := gin.New()
	return router, false
}

func checkNotFound(c *gin.Context) {
	c.Next()
	if c.Writer.Status() == 404 {
		c.JSON(404, gin.H{"message": "404 NOT FOUND"})
	}
}

func putMovie() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{"message": "PUT METHOD MOVIE"})
	}
}

func deleteMovie() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.IndentedJSON(204, gin.H{"message": "DELETE METHOD MOVIE"})
	}
}
