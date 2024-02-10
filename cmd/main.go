package main

import (
	"log"
	"net/http"

	"github.com/joaovds/rinha-crebito/configs"
	"github.com/joaovds/rinha-crebito/internal/infra/chi"
)

func main() {
	configs.LoadEnv()

  c := chi.SetupChi()

  log.Printf("Server running on port %s", configs.ENV.Port)
  log.Panic(http.ListenAndServe(":"+configs.ENV.Port, c.Mux))
}
