package routes

import (
	c "github.com/go-chi/chi/v5"
)

func SetupRoutes(mux *c.Mux) {
	handleClientRoutes(mux)
}
