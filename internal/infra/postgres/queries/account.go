package queries

const (
	GetAllAccounts      = "SELECT id, \"limit\", balance from accounts WHERE id = $1"
	GetAccountByID      = "SELECT id, \"limit\", balance from accounts WHERE id = $1"
	GetLastTransactions = "SELECT value, type_transaction, description, realized_at FROM transactions WHERE id = $1 ORDER BY realized_at DESC LIMIT 10"
)
