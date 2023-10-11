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
	router.DELETE("movies/:id", DeleteMovie())

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
