package address

import "github.com/csmistry/coin-tracker/backend/pkg/transaction"

type Address struct {
	Id           string                    `json:"id"`
	Balance      float64                   `json:"balance"`
	Transactions []transaction.Transaction `json:"transactions"`
	IsArchived   bool
}
