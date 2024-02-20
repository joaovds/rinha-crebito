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

func (r *AccountDBRepository) GetLastTransactions(accountId int) ([]*domain.LastTransaction, error) {
	transactions := make([]*domain.LastTransaction, 0)

	rows, err := r.db.Query(queries.GetLastTransactions, accountId)
	if err != nil {
		return transactions, err
	}

	for rows.Next() {
		var transaction domain.LastTransaction

		err = rows.Scan(&transaction.Value, &transaction.TypeTransaction, &transaction.Description, &transaction.RealizedAt)
		if err != nil {
			return transactions, err
		}

		transactions = append(transactions, &transaction)
	}

	return transactions, nil
}
