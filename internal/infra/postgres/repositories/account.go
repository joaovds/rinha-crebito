package repositories

import "github.com/joaovds/rinha-crebito/internal/domain"

type AccountDBRepository struct {}

var (
  NewAccountDBRepositoryInstance *AccountDBRepository
)

func NewAccountDBRepository() *AccountDBRepository {
  if NewAccountDBRepositoryInstance == nil {
    NewAccountDBRepositoryInstance = &AccountDBRepository{}
  }

  return NewAccountDBRepositoryInstance
}

func (r *AccountDBRepository) GetByID(id int) (*domain.Account, error) {
  return &domain.Account{
    ID: 1,
    Limit: 1000,
    Balance: 500,
  }, nil
}
