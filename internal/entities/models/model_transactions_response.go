package models

type TransactionsResponse struct {

	Transactions []Transaction `json:"transactions,omitempty"`
}
