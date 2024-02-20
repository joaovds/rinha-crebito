package queries

const (
	GetAllAccounts = "SELECT id, \"limit\", balance from accounts WHERE id = $1"
	GetAccountByID = "SELECT id, \"limit\", balance from accounts WHERE id = $1"
	UpdateAccountBalance  = "UPDATE accounts SET balance = $1 WHERE id = $2"
)
