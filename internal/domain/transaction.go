package domain

import "time"

type Transaction struct {
	ID              int
	Value           int    `json:"valor"`
	TypeTransaction string `json:"tipo"`
	Description     string `json:"descricao"`
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

type TransactionRepository interface {
	Create(transaction *Transaction) error
}

type TransactionUseCase interface {
	Create(transaction *Transaction) error
}
