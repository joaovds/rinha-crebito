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

func (c *ClientService) GetClientExtract(id int) (*dtos.ExtractDTO, error) {
	client, err := c.Crepo.GetClientByID(id)
	if err != nil {
		return nil, err
	}

	transactions, err := c.Crepo.GetLastTransactions(id)
	if err != nil {
		return nil, err
	}

	return &dtos.ExtractDTO{
		Balance: dtos.Balance{
			Total:       client.Balance,
			ExtractDate: time.Now(),
			Limit:       client.Limit,
		},
		LastTransactions: transactions,
	}, nil
}
