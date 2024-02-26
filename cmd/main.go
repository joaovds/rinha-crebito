package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/clientes/{id}/transacoes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if r.PathValue("id") == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, client " + r.PathValue("id")))
	})

	mux.HandleFunc("/clientes/{id}/extrato", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if r.PathValue("id") == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, client " + r.PathValue("id")))
	})

	log.Println("Starting server on http://localhost:9999")
	log.Fatal(http.ListenAndServe(":9999", mux))
}
