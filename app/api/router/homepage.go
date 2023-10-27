package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	// Use the Redirect method to redirect to the desired URL
	c.Redirect(http.StatusFound, "http://localhost:8080")
}
