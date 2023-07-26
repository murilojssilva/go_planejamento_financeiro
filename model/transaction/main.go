package transaction

import (
	"time"
)

// Transaction have transaction vars
type Transaction struct { // visibilidade pública por começar com letra maiúscula
	Title           string    `json:"title"`
	Amount          float32   `json:"amount"`
	TransactionType int       `json:"transaction_type"`
	CreatedAt       time.Time `json:"created_at"`
}

// Transactions is a Transaction collection
type Transactions []Transaction
