package handlers

import "github.com/joaovds/rinha-crebito/internal/domain"

func (r *CreateNewTransactionRequest) ToDomain(accountID int) *domain.Transaction {
	return domain.NewTransaction(0, r.Value, r.TypeTransaction, r.Description, accountID)
}
