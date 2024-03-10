package services

import (
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

	result, err := c.Crepo.InsertTransaction(&transaction)
	if err != nil {
		return nil, err
	}

	return &dtos.NewTransactionResponse{
		Balance: result.Balance,
		Limit:   result.Limit,
	}, nil
}
