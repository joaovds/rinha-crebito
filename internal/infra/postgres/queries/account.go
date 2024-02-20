package queries

const (
	GetAllAccounts = "SELECT id, \"limit\", balance FROM accounts"
	GetAccountByID = "SELECT id, \"limit\", balance FROM accounts WHERE id = $1"
	UpdateAccount  = "UPDATE accounts SET \"limit\" = $1, balance = $2 WHERE id = $3"
)
