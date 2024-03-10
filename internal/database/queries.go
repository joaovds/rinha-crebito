package database

const (
	GetClientByIDQuery = "SELECT id, \"limit\", balance from accounts WHERE id = $1"

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
	InsertTransactionQuery = "SELECT insert_transaction_and_update_client_balance($1, $2, $3, $4, $5)"
)
