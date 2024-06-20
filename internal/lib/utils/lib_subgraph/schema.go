package lib_subgraph

type DepositEvent struct {
	Id        string `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Amount    string `json:"amount"`
	Timestamp string `json:"timestamp"`
}

type WithdrawEvent struct {
	Id        string `json:"id"`
	To        string `json:"to"`
	Amount    string `json:"amount"`
	Nonce     int64  `json:"nonce"`
	Timestamp string `json:"timestamp"`
}
