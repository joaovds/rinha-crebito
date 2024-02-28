package services

import (
	"context"
	"time"

	"github.com/joaovds/rinha-crebito/internal/database"
	"github.com/joaovds/rinha-crebito/internal/dtos"
)

type ClientService struct {
	Crepo *database.Repository
}

func NewClientService() *ClientService {
	cr := database.NewRepository()

	return &ClientService{Crepo: cr}
}

func (c *ClientService) GetClientExtract(id int) (*dtos.ExtractResponse, error) {
	client, err := c.Crepo.GetClientByID(id)
	if err != nil {
		return nil, err
	}

	transactions, err := c.Crepo.GetLastTransactions(id)
	if err != nil {
		return nil, err
	}

	return &dtos.ExtractResponse{
		Balance: dtos.Balance{
			Total:       client.Balance,
			ExtractDate: time.Now(),
			Limit:       client.Limit,
		},
		LastTransactions: transactions,
	}, nil
}

func (c *ClientService) CreateNewTransaction(transaction dtos.CreateTransactionRequest) (*dtos.NewTransactionResponse, error) {
	isValidErr := transaction.IsValid()
	if isValidErr != nil {
		return nil, isValidErr
	}

	tx, err := c.Crepo.NewDBTransaction()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	client, err := c.Crepo.GetClientByIDForUpdate(tx, transaction.ClientID)
	if err != nil {
		return nil, err
	}

	var newBalance int

	if transaction.TypeTransaction == "d" {
		newBalance = client.Balance - transaction.Value

		if newBalance < -client.Limit {
			return nil, dtos.ErrIncosistentBalance
		}
	}

	if transaction.TypeTransaction == "c" {
		newBalance = client.Balance + transaction.Value
	}

	client.Balance = newBalance

	err = c.Crepo.InsertTransaction(tx, &transaction)
	if err != nil {
		return nil, err
	}

	err = c.Crepo.UpdateClientBalance(tx, client)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, err
	}

	return &dtos.NewTransactionResponse{
		Balance: newBalance,
		Limit:   client.Limit,
	}, nil
}
