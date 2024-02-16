package queries

const (
	GetLastTransactions = "SELECT id, value, type_transaction, description, realized_at, account_id FROM transactions ORDER BY realized_at DESC LIMIT 10"
)
