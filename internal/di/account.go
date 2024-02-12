package di

import (
	"github.com/joaovds/rinha-crebito/internal/infra/postgres/repositories"
	"github.com/joaovds/rinha-crebito/internal/usecases"
)

func NewAccountUsecases() *usecases.AccountUsecase {
  repo := repositories.NewAccountDBRepository()

  return usecases.NewAccountUsecase(*repo)
}
