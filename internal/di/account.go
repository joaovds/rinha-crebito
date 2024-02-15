package di

import (
	"github.com/joaovds/rinha-crebito/internal/data/usecases"
	"github.com/joaovds/rinha-crebito/internal/infra/postgres/repositories"
)

func NewAccountUsecases() *usecases.AccountUsecase {
	repo := repositories.NewAccountDBRepository()

	return usecases.NewAccountUsecase(*repo)
}
