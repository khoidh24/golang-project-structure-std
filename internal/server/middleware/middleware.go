package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PreventHomepageNoParams(c *gin.Context) {
	params := c.Request.URL.Query()
	if params.Get("name") == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Where's your name?"})
		return
	}
	c.Next()
}