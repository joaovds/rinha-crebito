package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/joaovds/rinha-crebito/internal/database"
	"github.com/joaovds/rinha-crebito/internal/services"
)

func GetClientExtract(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.PathValue("id") == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cserv := services.NewClientService()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	extract, err := cserv.GetClientExtract(id)
	if err != nil {
		if err == database.ErrClientNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}

	response, err := json.Marshal(extract)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.PathValue("id") == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transaction created for client " + r.PathValue("id")))
}
