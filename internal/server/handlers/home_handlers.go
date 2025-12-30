package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) HelloWorldHandler(c *gin.Context) {
	params := c.Request.URL.Query().Get("name")

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Hello " + params})
}
