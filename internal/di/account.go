package di

import (
	"sync"

	"github.com/joaovds/rinha-crebito/internal/data/usecases"
	"github.com/joaovds/rinha-crebito/internal/infra/postgres/repositories"
)

var (
	accountUsecase *usecases.AccountUsecase
	Once           sync.Once
)

func NewAccountUsecases() *usecases.AccountUsecase {
	Once.Do(func() {
		repo := repositories.NewAccountDBRepository()

		accountUsecase = usecases.NewAccountUsecase(repo)
	})

	return accountUsecase
}
