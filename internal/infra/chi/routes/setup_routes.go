package routes

import (
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(mux *chi.Mux) {
	handleClientRoutes(mux)
}
