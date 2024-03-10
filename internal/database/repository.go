package database

import (
	"context"
	"encoding/json"
	"errors"
	"log"

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

type InsertTransactionResponse struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

func NewInsertTransactionResponse(limit, balance int) *InsertTransactionResponse {
	return &InsertTransactionResponse{
		Limit:   limit,
		Balance: balance,
	}
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

func (r *Repository) InsertTransaction(transaction *dtos.CreateTransactionRequest) (*InsertTransactionResponse, error) {
	var insertResponseJSON string

	err := Db.QueryRow(
		context.Background(),
		InsertTransactionQuery,
		transaction.Value,
		transaction.TypeTransaction,
		transaction.Description,
		transaction.ClientID,
	).Scan(&insertResponseJSON)
	if err != nil {
		return nil, err
	}

	var insertResponse InsertTransactionResponse
	err = json.Unmarshal([]byte(insertResponseJSON), &insertResponse)
	if err != nil {
		return nil, err
	}

	return &insertResponse, nil
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
