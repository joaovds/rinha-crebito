package main

import (
	"log"
	"net/http"

	"github.com/joaovds/rinha-crebito/configs"
	"github.com/joaovds/rinha-crebito/internal/infra/chi"
	"github.com/joaovds/rinha-crebito/internal/infra/postgres"
)

func main() {
	configs.LoadEnv()

  conn := postgres.NewDatabase()
  defer conn.Pool.Close()

	c := chi.SetupChi()

	log.Printf("Server running on port %s", configs.ENV.Port)
	log.Panic(http.ListenAndServe(":"+configs.ENV.Port, c.Mux))
}
