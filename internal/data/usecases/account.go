package usecases

import (
	"time"

	"github.com/joaovds/rinha-crebito/internal/domain"
)

type AccountUsecase struct {
	repo domain.AccountRepository
}

var (
	NewAccountUsecaseInstance *AccountUsecase
)

func NewAccountUsecase(repo domain.AccountRepository) *AccountUsecase {
	if NewAccountUsecaseInstance == nil {
		NewAccountUsecaseInstance = &AccountUsecase{
			repo: repo,
		}
	}

	return NewAccountUsecaseInstance
}

func (a *AccountUsecase) GetAccountByID(id int) (*domain.Account, error) {
	return a.repo.GetByID(id)
}

func (a *AccountUsecase) GetExtract(id int) (*domain.Extract, error) {
	account, err := a.GetAccountByID(id)
	if err != nil {
		return &domain.Extract{}, err
	}

	transactions, err := a.repo.GetLastTransactions(id)
	if err != nil {
		return &domain.Extract{}, err
	}

	return &domain.Extract{
		Balance: domain.Balance{
			Total:       account.Balance,
			ExtractDate: time.Now(),
			Limit:       account.Limit,
		},
		LastTransactions: transactions,
	}, nil
}
