package dtos

import "time"

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

type Transaction struct {
	ID              int
	Value           int
	TypeTransaction string
	Description     string
	AccountID       int
	RealizedAt      time.Time
}

func NewTransaction(id, value int, typeTransaction, description string, accountID int) *Transaction {
	return &Transaction{
		ID:              id,
		Value:           value,
		TypeTransaction: typeTransaction,
		Description:     description,
		AccountID:       accountID,
		RealizedAt:      time.Now(),
	}
}

type ExtractDTO struct {
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
