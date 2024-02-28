package dtos

import (
	"errors"
	"time"
)

var (
	ErrIncosistentBalance = errors.New("inconsistent balance")
)

type Client struct {
	ID      int
	Limit   int
	Balance int
}

func NewClient(id, limit, balance int) *Client {
	return &Client{
		ID:      id,
		Limit:   limit,
		Balance: balance,
	}
}

type CreateTransactionRequest struct {
	Value           int    `json:"valor"`
	TypeTransaction string `json:"tipo"`
	Description     string `json:"descricao"`
	ClientID        int
}

func NewCreateTransactionRequest(id, value int, typeTransaction, description string, clientID int) *CreateTransactionRequest {
	return &CreateTransactionRequest{
		Value:           value,
		TypeTransaction: typeTransaction,
		Description:     description,
		ClientID:        clientID,
	}
}

func (ctr *CreateTransactionRequest) IsValid() error {
	if ctr.Value <= 0 {
		return ErrIncosistentBalance
	}
	if ctr.TypeTransaction != "d" && ctr.TypeTransaction != "c" {
		return ErrIncosistentBalance
	}
	if ctr.Description == "" || len(ctr.Description) > 10 {
		return ErrIncosistentBalance
	}

	return nil
}

type NewTransactionResponse struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

type ExtractResponse struct {
	Balance          Balance            `json:"saldo"`
	LastTransactions []*LastTransaction `json:"ultimas_transacoes"`
}

type Balance struct {
	Total       int       `json:"total"`
	ExtractDate time.Time `json:"data_extrato"`
	Limit       int       `json:"limite"`
}

type LastTransaction struct {
	Value           int       `json:"valor"`
	TypeTransaction string    `json:"tipo"`
	Description     string    `json:"descricao"`
	RealizedAt      time.Time `json:"realizada_em"`
}
