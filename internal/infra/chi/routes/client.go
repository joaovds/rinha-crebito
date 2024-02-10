package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaovds/rinha-crebito/internal/infra/chi/handlers"
)

func handleClientRoutes(mux *chi.Mux) {
	clientHandlers := handlers.NewClientHandler()

	mux.Route("/clientes", func(router chi.Router) {
		router.Get("/{id}/extrato", func(w http.ResponseWriter, r *http.Request) {
      clientHandlers.GetExtract(w, r)
		})
	})
}
