package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/joaovds/rinha-crebito/internal/database"
	"github.com/joaovds/rinha-crebito/internal/dtos"
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

		w.WriteHeader(http.StatusInternalServerError)
		return
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

	var requestData dtos.CreateTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "failed to decode request body", http.StatusUnprocessableEntity)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	requestData.ClientID = id

	cserv := services.NewClientService()

	newTransactionResult, err := cserv.CreateNewTransaction(requestData)
	if err != nil {
		if err == database.ErrClientNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err == dtos.ErrIncosistentBalance {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(newTransactionResult)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
