package models

import (
	"time"
)

type Receipt struct {

	Id int64 `json:"id"`

	UserAddress string `json:"user_address"`

	GatewayAddress string `json:"gateway_address"`

	FinishedAt time.Time `json:"finished_at,omitempty"`

	Cost string `json:"cost,omitempty"`

	Proof string `json:"proof,omitempty"`

	TxnId string `json:"txn_id,omitempty"`

	Status string `json:"status,omitempty"`
}
