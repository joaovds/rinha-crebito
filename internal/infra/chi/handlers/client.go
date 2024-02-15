package handlers

import (
	"encoding/json"
	"net/http"

	cc "github.com/joaovds/rinha-crebito/internal/infra/chi/contracts"
	"github.com/joaovds/rinha-crebito/internal/infra/chi/middlewares"
)

type clientHandler struct{}

func NewClientHandler() *clientHandler {
	return &clientHandler{}
}

func (h *clientHandler) GetExtract(w http.ResponseWriter, r *http.Request) {
	account, ok := middlewares.AccountFromContext(r.Context())
	if !ok {
		http.Error(w, cc.NewErrorResponse(http.StatusNotFound, "account not found").String(), http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(account)

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *clientHandler) CreateNewTransaction(w http.ResponseWriter, r *http.Request) {
	account, ok := middlewares.AccountFromContext(r.Context())
	if !ok {
		http.Error(w, cc.NewErrorResponse(http.StatusNotFound, "account not found").String(), http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(account)

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
