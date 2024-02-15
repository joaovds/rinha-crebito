package domain

import "errors"

var (
	ErrAccountNotFound = errors.New("account not found")
)

type Account struct {
	ID      int
	Limit   int
	Balance int
}

func NewAccount(id, limit, balance int) *Account {
	return &Account{
		ID:      id,
		Limit:   limit,
		Balance: balance,
	}
}

type AccountRepository interface {
	GetByID(id int) (*Account, error)
}

type AccountUseCase interface {
	GetByID(id int) (*Account, error)
}
