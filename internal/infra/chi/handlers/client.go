package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joaovds/rinha-crebito/internal/di"
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

	accountUC := di.NewAccountUsecases()
	extract, err := accountUC.GetExtract(account.ID)
	if err != nil {
		http.Error(w, cc.NewErrorResponse(http.StatusInternalServerError, "internal server error").String(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	response, _ := json.Marshal(extract)

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
