package main

import (
	"log"
	"net/http"

	"github.com/joaovds/rinha-crebito/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/clientes/{id}/transacoes", handlers.CreateTransaction)

	mux.HandleFunc("/clientes/{id}/extrato", handlers.GetClientExtract)

	log.Println("Starting server on http://localhost:9999")
	log.Fatal(http.ListenAndServe(":9999", mux))
}
