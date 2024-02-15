package repositories

import (
	"database/sql"

	"github.com/joaovds/rinha-crebito/internal/domain"
	"github.com/joaovds/rinha-crebito/internal/infra/postgres"
)

type AccountDBRepository struct {
	db *sql.DB
}

var (
	NewAccountDBRepositoryInstance *AccountDBRepository
)

func NewAccountDBRepository() *AccountDBRepository {
	db, _ := postgres.GetConnection()

	if NewAccountDBRepositoryInstance == nil {
		NewAccountDBRepositoryInstance = &AccountDBRepository{
			db: db,
		}
	}

	return NewAccountDBRepositoryInstance
}

func (r *AccountDBRepository) GetByID(id int) (*domain.Account, error) {
	var account domain.Account

	err := r.db.QueryRow(
		"SELECT id, \"limit\", balance from clients WHERE id = $1",
		id,
	).Scan(&account.ID, &account.Limit, &account.Balance)
	if err != nil {
		println(err.Error())
		return &domain.Account{}, err
	}

	return &account, nil
}
