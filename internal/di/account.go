package di

import (
	"sync"

	"github.com/joaovds/rinha-crebito/internal/data/usecases"
	"github.com/joaovds/rinha-crebito/internal/infra/postgres/repositories"
)

var (
	accountUsecase     *usecases.AccountUsecase
	transactionUsecase *usecases.TransactionUsecase
	OnceAUC            sync.Once
	OnceTUC            sync.Once
)

func NewAccountUsecases() *usecases.AccountUsecase {
	OnceAUC.Do(func() {
		repo := repositories.NewAccountDBRepository()

		accountUsecase = usecases.NewAccountUsecase(repo)
	})

	return accountUsecase
}

func NewTransactionUsecases() *usecases.TransactionUsecase {
	OnceTUC.Do(func() {
		repo := repositories.NewTransactionDBRepository()

		transactionUsecase = usecases.NewTransactionUsecase(repo)
	})

	return transactionUsecase
}
