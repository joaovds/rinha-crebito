package domain

import "time"

type Extract struct {
	Balance          Balance            `json:"saldo"`
	LastTransactions []*LastTransaction `json:"ultimas_transacoes"`
}

type Balance struct {
	Total       int       `json:"total"`
	ExtractDate time.Time `json:"data_extrato"`
	Limit       int       `json:"limite"`
}

type LastTransaction struct {
	Value           int       `json:"valor"`
	TypeTransaction string    `json:"tipo"`
	Description     string    `json:"descricao"`
	RealizedAt      time.Time `json:"realizada_em"`
}
