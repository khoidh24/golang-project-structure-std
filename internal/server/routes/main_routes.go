package routes

import (
	"rushly/internal/server/handlers"
	v1 "rushly/internal/server/routes/v1"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all routes for the server.
// It groups versioned routes (e.g., /v1) and delegates to specific route handlers.
func RegisterRoutes(r *gin.Engine, handlers *handlers.Handler) {
	// Group all v1 routes
	v1Group := r.Group("/v1")
	{
		// Pass the handlers instance effectively, so sub-routes can pick what they need
		v1.RegisterHealthRoutes(v1Group, handlers)
		// Example: v1.RegisterNoteRoutes(v1Group, h)
		v1.RegisterHomeRoutes(v1Group, handlers)
	}
}
