package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handleClientRoutes(mux *chi.Mux) {
	mux.Route("/clientes", func(router chi.Router) {
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte("Hello, World!"))
		})
	})
}
