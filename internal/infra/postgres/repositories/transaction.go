package repositories

import (
	"database/sql"
	"time"

	"github.com/joaovds/rinha-crebito/internal/domain"
	"github.com/joaovds/rinha-crebito/internal/infra/postgres"
	"github.com/joaovds/rinha-crebito/internal/infra/postgres/queries"
)

type TransactionDBRepository struct {
	db *sql.DB
}

var (
	NewTransactionDBRepositoryInstance *TransactionDBRepository
)

func NewTransactionDBRepository() *TransactionDBRepository {
	db, _ := postgres.GetConnection()

	if NewTransactionDBRepositoryInstance == nil {
		NewTransactionDBRepositoryInstance = &TransactionDBRepository{
			db: db,
		}
	}

	return NewTransactionDBRepositoryInstance
}

func (r *TransactionDBRepository) Create(transaction *domain.Transaction) error {
	_, err := r.db.Exec(
		queries.InsertTransaction,
		transaction.Value,
		transaction.TypeTransaction,
		transaction.Description,
		time.Now(),
		transaction.AccountID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionDBRepository) UpdateAccountBalance(account *domain.Account) error {
	_, err := t.db.Exec(
		queries.UpdateAccountBalance,
		account.Balance,
		account.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
