package models

import (
	"time"
)

type CreateReceiptRequest struct {

	UserAddress string `json:"user_address"`

	GatewayAddress string `json:"gateway_address"`

	FinishedAt time.Time `json:"finished_at"`

	Cost string `json:"cost"`

	Proof string `json:"proof"`

	TxnId string `json:"txn_id"`

	Status string `json:"status"`
}
