package router

import (
	"hz/config"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func ApiHandler() {
	router, done := routerSettings()
	if done {
		return
	}

	router.Use(CORSMiddleware())

	APIRequest(router)

	startRouter(router)
}

func APIRequest(router *gin.Engine) {
	router.Use(checkNotFound)
	router.GET("/", Homepage)
	router.GET("/movies", getMovies())
	router.GET("/information", getInformation())
	router.GET("/movies/:id", GetMovieDetails())
	router.POST("/movies", PostMovies())
	router.DELETE("movies/:id", DeleteMovie())
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

// CORSMiddleware sets up CORS headers to allow requests from any origin.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
