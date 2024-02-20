package queries

const (
	GetLastTransactions = "SELECT value, type_transaction, description, realized_at FROM transactions WHERE id = $1 ORDER BY realized_at DESC LIMIT 10"
	InsertTransaction   = "INSERT INTO transactions (value, type_transaction, description, realized_at, account_id) VALUES ($1, $2, $3, $4, $5)"
)
