package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
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
	router.POST("/movies", postMovies())
	router.PUT("/movies/:id", putMovie())
	router.DELETE("movies/:id", deleteMovie())

	startRouter(router)
}

func startRouter(router *gin.Engine) {
	errorRouter := router.Run(":8090")
	if errorRouter != nil {
		return
	}
}

func routerSettings() (*gin.Engine, bool) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.ForwardedByClientIP = true
	errorTrustedProxies := router.SetTrustedProxies([]string{"127.0.0.1", "0.0.0.0"})
	if errorTrustedProxies != nil {
		return nil, true
	}
	return router, false
}

func checkNotFound(c *gin.Context) {
	c.Next()

	if c.Writer.Status() == 404 {
		c.JSON(404, gin.H{"message": "404 NOT FOUND"})
	}
}

func postMovies() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(201, gin.H{"message": "POST METHOD MOVIES"})
	}
}

func putMovie() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "PUT METHOD MOVIE"})
	}
}

func deleteMovie() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(204, gin.H{"message": "DELETE METHOD MOVIE"})
	}
}
