package main

import (
	"log"
	"net/http"

	"github.com/joaovds/rinha-crebito/internal/database"
	"github.com/joaovds/rinha-crebito/internal/handlers"
)

func main() {
	db, _ := database.SetupDatabase()
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/clientes/{id}/transacoes", handlers.CreateTransaction)

	mux.HandleFunc("/clientes/{id}/extrato", handlers.GetClientExtract)

	log.Println("Starting server on http://localhost:9999")
	log.Fatal(http.ListenAndServe(":9999", mux))
}

func startDatabase() {
	db, _ := database.SetupDatabase()
	defer db.Close()
}
