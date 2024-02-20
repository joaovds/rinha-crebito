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

func (a *TransactionUsecase) Create(transaction *domain.Transaction) error {
	a.repo.Create(transaction)

	return nil
}
