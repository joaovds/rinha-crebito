package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joaovds/rinha-crebito/internal/database"
	"github.com/joaovds/rinha-crebito/internal/handlers"
)

func main() {
	db, _ := database.SetupDatabase()
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /clientes/{id}/transacoes", handlers.CreateTransaction)

	mux.HandleFunc("GET /clientes/{id}/extrato", handlers.GetClientExtract)

	APIPORT := os.Getenv("PORT")
	if APIPORT == "" {
		APIPORT = "9999"
	}

	log.Println(fmt.Sprintf("Starting server on http://localhost:%s", APIPORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", APIPORT), mux))
}

func startDatabase() {
	db, _ := database.SetupDatabase()
	defer db.Close()
}
