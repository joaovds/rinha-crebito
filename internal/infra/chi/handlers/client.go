package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	c "github.com/go-chi/chi/v5"
	"github.com/joaovds/rinha-crebito/internal/di"
)

type clientHandler struct{}

func NewClientHandler() *clientHandler {
	return &clientHandler{}
}

func (h *clientHandler) GetExtract(w http.ResponseWriter, r *http.Request) {
	clientIdParam := c.URLParam(r, "id")
	clientId, err := strconv.Atoi(clientIdParam)
	if err != nil {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	accountUC := di.NewAccountUsecases()

	account, err := accountUC.GetAccountByID(clientId)
	if err != nil {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(account)

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
