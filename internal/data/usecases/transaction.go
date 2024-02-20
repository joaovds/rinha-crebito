package usecases

import (
	"github.com/joaovds/rinha-crebito/internal/domain"
)

type TransactionUsecase struct {
	repo domain.TransactionRepository
}

var (
	NewTransactionUsecaseInstance *TransactionUsecase
)

func NewTransactionUsecase(repo domain.TransactionRepository) *TransactionUsecase {
	if NewTransactionUsecaseInstance == nil {
		NewTransactionUsecaseInstance = &TransactionUsecase{
			repo: repo,
		}
	}

	return NewTransactionUsecaseInstance
}

func (a *TransactionUsecase) Create(transaction *domain.Transaction, account *domain.Account) error {
  err := transaction.Validate()
  if err != nil {
    return err
  }

  var newBalance int

  if transaction.TypeTransaction == "d" {
    newBalance = account.Balance - transaction.Value

    if newBalance < -account.Limit {
      return domain.ErrIncosistentBalance
    }
  }

  if transaction.TypeTransaction == "c" {
    newBalance = account.Balance + transaction.Value
  }

  account.Balance = newBalance

  err = a.repo.Create(transaction)
  if err != nil {
    return err
  }

  err = a.repo.UpdateAccountBalance(account)
  if err != nil {
    return err
  }

	return nil
}
