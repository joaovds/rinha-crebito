package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joaovds/rinha-crebito/internal/di"
	"github.com/joaovds/rinha-crebito/internal/domain"
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

	var requestBody CreateNewTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, cc.NewErrorResponse(http.StatusUnprocessableEntity, "failed to decode request body").String(), http.StatusUnprocessableEntity)
		return
	}

	transactionUC := di.NewTransactionUsecases()
	err := transactionUC.Create(requestBody.ToDomain(account.ID), account)
	if err != nil {
		if err == domain.ErrIncosistentBalance {
			http.Error(w, cc.NewErrorResponse(http.StatusUnprocessableEntity, "transaction would exceed account limit").String(), http.StatusUnprocessableEntity)
			return
		}

		http.Error(w, cc.NewErrorResponse(http.StatusUnprocessableEntity, "internal server error").String(), http.StatusUnprocessableEntity)
		log.Println(err.Error())
		return
	}

	response, _ := json.Marshal(NewCreateNewTransactionResponse(account.Limit, account.Balance))

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
