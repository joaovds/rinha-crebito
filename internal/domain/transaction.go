package domain

import (
	"errors"
	"time"
)

var (
	ErrIncosistentBalance = errors.New("inconsistent balance")
)

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

func (t *Transaction) Validate() error {
	if t.Value <= 0 {
		return ErrIncosistentBalance
	}
	if t.TypeTransaction != "d" && t.TypeTransaction != "c" {
		return ErrIncosistentBalance
	}
	if t.Description == "" || len(t.Description) > 10 {
		return ErrIncosistentBalance
	}

	return nil
}

type TransactionRepository interface {
	Create(transaction *Transaction) error
	UpdateAccountBalance(account *Account) error
}

type TransactionUseCase interface {
	Create(transaction *Transaction) error
}
