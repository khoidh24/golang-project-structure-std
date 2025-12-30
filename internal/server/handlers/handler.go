package handlers

import "rushly/internal/database"

type Handler struct {
	db database.Service
}

func New(db database.Service) *Handler {
	return &Handler{
		db: db,
	}
}
