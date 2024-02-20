package queries

const (
	GetLastTransactions = "SELECT id, value, type_transaction, description, realized_at, account_id FROM transactions ORDER BY realized_at DESC LIMIT 10"
	InsertTransaction   = "INSERT INTO transactions (value, type_transaction, description, realized_at, account_id) VALUES ($1, $2, $3, $4, $5)"
)
