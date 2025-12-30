package v1

import (
	"rushly/internal/server/handlers"
	"rushly/internal/server/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterHomeRoutes(r *gin.RouterGroup, handlers *handlers.Handler) {
	r.GET("/", middleware.PreventHomepageNoParams, handlers.HelloWorldHandler)
}
