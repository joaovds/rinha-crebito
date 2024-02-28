package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joaovds/rinha-crebito/internal/dtos"
)

var (
	ErrClientNotFound = errors.New("client not found")
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) NewDBTransaction() (pgx.Tx, error) {
	tx, err := Db.Begin(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return tx, nil
}

func (r *Repository) GetClientByID(id int) (*dtos.Client, error) {
	var client dtos.Client

	err := Db.QueryRow(
		context.Background(),
		GetClientByIDQuery,
		id,
	).Scan(&client.ID, &client.Limit, &client.Balance)
	if err != nil {
		log.Println(err)

		if err == pgx.ErrNoRows {
			return &dtos.Client{}, ErrClientNotFound
		}

		return &dtos.Client{}, err
	}

	return &client, nil
}

func (r *Repository) GetClientByIDForUpdate(tx pgx.Tx, id int) (*dtos.Client, error) {
	var client dtos.Client

	err := tx.QueryRow(
		context.Background(),
		GetClientByIDForUpdateQuery,
		id,
	).Scan(&client.ID, &client.Limit, &client.Balance)
	if err != nil {
		log.Println(err)

		if err == pgx.ErrNoRows {
			return &dtos.Client{}, ErrClientNotFound
		}

		return &dtos.Client{}, err
	}

	return &client, nil
}

func (r *Repository) UpdateClientBalance(tx pgx.Tx, client *dtos.Client) error {
	_, err := tx.Exec(
		context.Background(),
		UpdateClientBalanceQuery,
		client.Balance,
		client.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) InsertTransaction(tx pgx.Tx, transaction *dtos.CreateTransactionRequest) error {
	_, err := tx.Exec(
		context.Background(),
		InsertTransactionQuery,
		transaction.Value,
		transaction.TypeTransaction,
		transaction.Description,
		time.Now(),
		transaction.ClientID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetLastTransactions(clientId int) ([]*dtos.LastTransaction, error) {
	transactions := make([]*dtos.LastTransaction, 0)

	rows, err := Db.Query(context.Background(), GetLastTransactionsQuery, clientId)
	if err != nil {
		log.Println(err)

		return transactions, err
	}

	for rows.Next() {
		var transaction dtos.LastTransaction

		err = rows.Scan(&transaction.Value, &transaction.TypeTransaction, &transaction.Description, &transaction.RealizedAt)
		if err != nil {
			log.Println(err)

			return transactions, err
		}

		transactions = append(transactions, &transaction)
	}

	return transactions, nil
}
