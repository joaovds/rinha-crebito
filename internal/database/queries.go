package database

const (
	GetAllClientsQuery          = "SELECT id, \"limit\", balance from accounts WHERE id = $1"
	GetClientByIDQuery          = "SELECT id, \"limit\", balance from accounts WHERE id = $1"
	GetClientByIDForUpdateQuery = "SELECT id, \"limit\", balance from accounts WHERE id = $1 FOR UPDATE"
	UpdateClientBalanceQuery    = "UPDATE accounts SET balance = $1 WHERE id = $2"

	GetLastTransactionsQuery = `
  SELECT
    value,
    type_transaction,
    description,
    realized_at
  FROM
    transactions tran
  JOIN accounts ac ON
    tran.account_id = ac.id
  WHERE
    ac.id = $1
  ORDER BY
    realized_at DESC
  LIMIT 10
  `
	InsertTransactionQuery = "INSERT INTO transactions (value, type_transaction, description, realized_at, account_id) VALUES ($1, $2, $3, $4, $5)"
)
