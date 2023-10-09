package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func ApiHandler() {

	router := gin.Default()
	router.GET("/movies", GetMovies())

	router.Run(":8090")
}
