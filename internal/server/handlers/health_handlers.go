package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, handler.db.Health())
}
