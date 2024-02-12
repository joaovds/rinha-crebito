package usecases

import (
	"github.com/joaovds/rinha-crebito/internal/domain"
	"github.com/joaovds/rinha-crebito/internal/infra/postgres/repositories"
)

type AccountUsecase struct {
  repo repositories.AccountDBRepository
}

var (
  NewAccountUsecaseInstance *AccountUsecase
)

func NewAccountUsecase(repo repositories.AccountDBRepository) *AccountUsecase {
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
