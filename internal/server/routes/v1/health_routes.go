package v1

import (
	"rushly/internal/server/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(r *gin.RouterGroup, handlers *handlers.Handler) {
	r.GET("/health", handlers.HealthHandler)
}
