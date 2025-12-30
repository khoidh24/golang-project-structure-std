package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"rushly/internal/server/handlers"
	"rushly/internal/server/routes"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	// Initialize handlers with dependencies
	handlers := handlers.New(s.db)

	// Delegate route registration to the routes package.
	routes.RegisterRoutes(router, handlers)

	return router
}
