package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaovds/rinha-crebito/internal/infra/chi/handlers"
	"github.com/joaovds/rinha-crebito/internal/infra/chi/middlewares"
)

func handleClientRoutes(mux *chi.Mux) {
	clientHandlers := handlers.NewClientHandler()

	mux.Route("/clientes", func(router chi.Router) {
		router.With(middlewares.CheckIDParam).Post("/{id}/transacoes", func(w http.ResponseWriter, r *http.Request) {
			clientHandlers.CreateNewTransaction(w, r)
		})

		router.With(middlewares.CheckIDParam).Get("/{id}/extrato", func(w http.ResponseWriter, r *http.Request) {
			clientHandlers.GetExtract(w, r)
		})
	})
}
