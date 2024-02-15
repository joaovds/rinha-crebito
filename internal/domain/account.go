package domain

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
