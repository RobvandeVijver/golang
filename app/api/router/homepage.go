package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	redirectURL := "http://127.0.0.1:5500/svelte-app/public/index.html"
	c.Redirect(http.StatusFound, redirectURL)
}
