package handlers

type CreateNewTransactionRequest struct {
	Value           int    `json:"valor"`
	TypeTransaction string `json:"tipo"`
	Description     string `json:"descricao"`
}

func NewCreateNewTransactionRequest(value int, typeTransaction, description string) *CreateNewTransactionRequest {
  return &CreateNewTransactionRequest{
    Value:           value,
    TypeTransaction: typeTransaction,
    Description:     description,
  }
}

type CreateNewTransactionResponse struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

func NewCreateNewTransactionResponse(limit, balance int) *CreateNewTransactionResponse {
	return &CreateNewTransactionResponse{
		Limit:   limit,
		Balance: balance,
	}
}
