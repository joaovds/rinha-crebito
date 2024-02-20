package repositories

import (
	"database/sql"

	"github.com/joaovds/rinha-crebito/internal/domain"
	"github.com/joaovds/rinha-crebito/internal/infra/postgres"
	"github.com/joaovds/rinha-crebito/internal/infra/postgres/queries"
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
		queries.GetAccountByID,
		id,
	).Scan(&account.ID, &account.Limit, &account.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return &domain.Account{}, domain.ErrAccountNotFound
		}

		return &domain.Account{}, err
	}

	return &account, nil
}

func (r *AccountDBRepository) Update(account *domain.Account) error {
	_, err := r.db.Exec(
		queries.UpdateAccount,
		account.Limit,
		account.Balance,
		account.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
