package usecases

import (
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
