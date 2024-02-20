package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joaovds/rinha-crebito/internal/domain"
	cc "github.com/joaovds/rinha-crebito/internal/infra/chi/contracts"
	"github.com/joaovds/rinha-crebito/internal/infra/chi/middlewares"
	"github.com/joaovds/rinha-crebito/internal/infra/postgres/repositories"
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

type CreateNewTransactionResponse struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

func NewCreateNewTransactionResponse(limit, balance int) *CreateNewTransactionResponse {
	return &CreateNewTransactionResponse{
		Limit:   limit,
		Balance: balance,
	}
}

func (h *clientHandler) CreateNewTransaction(w http.ResponseWriter, r *http.Request) {
	account, ok := middlewares.AccountFromContext(r.Context())
	if !ok {
		http.Error(w, cc.NewErrorResponse(http.StatusNotFound, "account not found").String(), http.StatusNotFound)
		return
	}

	var transaction domain.Transaction

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}
	transaction.AccountID = account.ID

	if transaction.TypeTransaction == "d" {
		if account.Balance-transaction.Value < -account.Limit {
			http.Error(w, "transaction would exceed account limit", http.StatusUnprocessableEntity)
			return
		}
		account.Balance = account.Balance - transaction.Value
	} else if transaction.TypeTransaction == "c" {
		account.Limit = account.Limit - transaction.Value
	}

	err := repositories.NewTransactionDBRepository().CreateTransaction(transaction)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create transaction: %v", err), http.StatusBadRequest)
		return
	}

	err = repositories.NewAccountDBRepository().Update(account)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to update account: %v", err), http.StatusBadRequest)
		return
	}

	accountNew := NewCreateNewTransactionResponse(account.Limit, account.Balance)
	response, _ := json.Marshal(accountNew)
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
